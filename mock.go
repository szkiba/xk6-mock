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

package mock

import (
	"context"
	"net"
	"net/http"
	"net/url"

	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"github.com/szkiba/xk6-mock/muxpress"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/lib"
)

type GlobalMock struct{}

func NewGlobalMock() *GlobalMock {
	return new(GlobalMock)
}

func (g *GlobalMock) NewModuleInstancePerVU() interface{} {
	l := logrus.StandardLogger().Error
	server := muxpress.NewServer(l)

	return &Mock{
		server:      server,
		errorLogger: l,
	}
}

type Mock struct {
	server      *muxpress.Server
	errorLogger func(...interface{})
}

func (m *Mock) XServer(ctxPtr *context.Context) (interface{}, error) {
	errorLogger := m.errorLogger

	if state := lib.GetState(*ctxPtr); state != nil {
		errorLogger = state.Logger.Error
	}

	rt := common.GetRuntime(*ctxPtr)
	server := muxpress.NewServer(errorLogger)

	return common.Bind(rt, server, ctxPtr), nil
}

func (m *Mock) requireBind(ctx context.Context) {
	rt := common.GetRuntime(ctx)

	if v := rt.GlobalObject().Get(mockModuleGlobalKey); v != nil && !goja.IsUndefined(v) && !goja.IsNull(v) {
		return
	}

	common.BindToGlobal(rt, map[string]interface{}{mockModuleGlobalKey: m})
}

func (m *Mock) requireListen() error {
	if m.server.IsRunning() {
		return nil
	}

	_, err := m.server.Start(0, "")

	return err
}

func (m *Mock) Addr() (*net.TCPAddr, error) {
	if err := m.requireListen(); err != nil {
		return nil, err
	}

	return m.server.Addr(), nil
}

func (m *Mock) Resolve(orig string) (string, error) {
	if err := m.requireListen(); err != nil {
		return orig, err
	}

	addr := m.server.Addr().String()

	if orig == "" {
		return "http://" + addr, nil
	}

	u, err := url.Parse(orig)
	if err != nil {
		return "", err
	}

	if u.Hostname() == "127.0.0.1" {
		return orig, nil
	}

	u.Host = addr
	u.Scheme = "http"

	return u.String(), nil
}

func (m *Mock) Handle(ctx context.Context, method string, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	if err := m.requireListen(); err != nil {
		return nil, err
	}

	m.requireBind(ctx)

	m.server.Handle(method, path, callback...)

	return m, nil
}

func (m *Mock) Use(ctx context.Context, callback ...muxpress.CallbackFunc) (*Mock, error) {
	if err := m.requireListen(); err != nil {
		return nil, err
	}

	m.requireBind(ctx)

	m.server.Use(callback...)

	return m, nil
}

func (m *Mock) All(ctx context.Context, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	return m.Handle(ctx, "", path, callback...)
}

func (m *Mock) Get(ctx context.Context, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	return m.Handle(ctx, http.MethodGet, path, callback...)
}

func (m *Mock) Head(ctx context.Context, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	return m.Handle(ctx, http.MethodHead, path, callback...)
}

func (m *Mock) Post(ctx context.Context, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	return m.Handle(ctx, http.MethodPost, path, callback...)
}

func (m *Mock) Put(ctx context.Context, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	return m.Handle(ctx, http.MethodPut, path, callback...)
}

func (m *Mock) Patch(ctx context.Context, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	return m.Handle(ctx, http.MethodPatch, path, callback...)
}

func (m *Mock) Delete(ctx context.Context, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	return m.Handle(ctx, http.MethodDelete, path, callback...)
}

func (m *Mock) Options(ctx context.Context, path string, callback ...muxpress.CallbackFunc) (*Mock, error) {
	return m.Handle(ctx, http.MethodOptions, path, callback...)
}

const mockModuleGlobalKey = "__xk6_mock_module"
