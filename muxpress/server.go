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
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var _ http.Handler = new(Server)

type Server struct {
	Router      *mux.Router
	listener    net.Listener
	errorLogger func(...interface{})
	middlewares callbackChain
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
	return &Server{
		errorLogger: log,
		Router:      mux.NewRouter(),
		middlewares: make(callbackChain, 0),
		listener:    nil,
	}
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

func (s *Server) Listen(port int, host string) (*Server, error) {
	if err := s.listen(port, host); err != nil {
		return nil, err
	}

	return s, http.Serve(s.listener, s.Router)
}

func (s *Server) Start(port int, host string) (*Server, error) {
	if err := s.listen(port, host); err != nil {
		return nil, err
	}

	go func() {
		if err := http.Serve(s.listener, s.Router); err != nil {
			s.errorLogger(err)
		}
	}()

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
		defer func() {
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
