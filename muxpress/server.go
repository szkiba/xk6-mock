// MIT License
//
// Copyright (c) 2021 Iván Szkiba
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package muxpress

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

var _ http.Handler = new(Server)

type Server struct {
	Router      *mux.Router
	listener    net.Listener
	errorLogger func(...interface{})
	middlewares callbackChain
	srv         *http.Server
	running     uint32
	mu          sync.Mutex
}

type CallbackFunc func(req *Request, res *Response, next func())

type callbackChain []CallbackFunc

func (h callbackChain) call(req *Request, res *Response) bool {
	current := 0
	cont := true
	call := func() {
		f := h[current]
		current++

		cont = false

		f(req, res, func() {
			cont = true
		})
	}

	for cont && current < len(h) {
		call()
	}

	return cont
}

func NewServer(log func(...interface{})) *Server {
	s := &Server{
		errorLogger: log,
		Router:      mux.NewRouter(),
		middlewares: make(callbackChain, 0),
		listener:    nil,
		srv:         nil,
		running:     0,
	}

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) listen(port int, host string) (err error) {
	if host == "" {
		host = "127.0.0.1"
	}

	s.listener, err = net.Listen("tcp", addr(host, port))
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Addr() *net.TCPAddr {
	if s.listener == nil {
		return nil
	}

	addr, _ := s.listener.Addr().(*net.TCPAddr)

	return addr
}

func (s *Server) IsRunning() bool {
	return atomic.LoadUint32(&s.running) != 0
}

func (s *Server) Listen(port int, host string) (*Server, error) {
	if err := s.listen(port, host); err != nil {
		return nil, err
	}

	return s, s.run()
}

func (s *Server) run() error {
	s.srv = new(http.Server)
	s.srv.Handler = s

	atomic.StoreUint32(&s.running, 1)

	err := s.srv.Serve(s.listener)

	if errors.Is(err, http.ErrServerClosed) {
		err = nil
	}

	atomic.StoreUint32(&s.running, 0)

	return err
}

func (s *Server) Start(port int, host string) (*Server, error) {
	if err := s.listen(port, host); err != nil {
		return nil, err
	}

	done := make(chan bool)

	go func(start chan bool) {
		start <- true

		if err := s.run(); err != nil {
			s.errorLogger(err)
		}
	}(done)

	<-done

	return s, nil
}

func (s *Server) Stop(timeout int) (*Server, error) {
	if !s.IsRunning() {
		return nil, ErrServerNotRunning
	}

	if timeout == 0 {
		timeout = shutdownPeriod
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) Use(callback ...CallbackFunc) {
	s.middlewares = append(s.middlewares, callback...)
}

func (s *Server) Handle(method string, path string, callback ...CallbackFunc) *Server {
	var r *mux.Route

	if len(method) == 0 {
		r = s.Router.NewRoute()
	} else {
		r = s.Router.Methods(method)
	}

	r.Path(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		defer func() {
			s.mu.Unlock()
			if r := recover(); r != nil {
				s.errorLogger(r)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		req, err := newRequest(r)
		if err != nil {
			s.errorLogger(err)

			return
		}

		res := newResponse(w)

		if s.middlewares.call(req, res) {
			callbackChain(callback).call(req, res)
		}
	})

	return s
}

func addr(host string, port int) string {
	return host + ":" + strconv.Itoa(port)
}

func (s *Server) All(path string, callback ...CallbackFunc) *Server {
	return s.Handle("", path, callback...)
}

func (s *Server) Get(path string, callback ...CallbackFunc) *Server {
	return s.Handle(http.MethodGet, path, callback...)
}

func (s *Server) Head(path string, callback ...CallbackFunc) *Server {
	return s.Handle(http.MethodHead, path, callback...)
}

func (s *Server) Post(path string, callback ...CallbackFunc) *Server {
	return s.Handle(http.MethodPost, path, callback...)
}

func (s *Server) Put(path string, callback ...CallbackFunc) *Server {
	return s.Handle(http.MethodPut, path, callback...)
}

func (s *Server) Patch(path string, callback ...CallbackFunc) *Server {
	return s.Handle(http.MethodPatch, path, callback...)
}

func (s *Server) Delete(path string, callback ...CallbackFunc) *Server {
	return s.Handle(http.MethodDelete, path, callback...)
}

func (s *Server) Options(path string, callback ...CallbackFunc) *Server {
	return s.Handle(http.MethodOptions, path, callback...)
}

var ErrServerNotRunning = errors.New("Server not running")

const shutdownPeriod = 1 // seconds
