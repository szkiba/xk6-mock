package mock

import (
	"context"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules/k6/http"
	"go.k6.io/k6/lib/netext/httpext"
)

type GlobalHTTP struct {
	*http.GlobalHTTP
}

func NewGlobalHTTP() *GlobalHTTP {
	return new(GlobalHTTP)
}

func (g *GlobalHTTP) NewModuleInstancePerVU() interface{} {
	return &HTTP{HTTP: g.GlobalHTTP.NewModuleInstancePerVU().(*http.HTTP)}
}

type HTTP struct {
	*http.HTTP
}

func (h *HTTP) Request(ctx context.Context, method string, url goja.Value, args ...goja.Value) (*http.Response, error) {
	rt := common.GetRuntime(ctx)

	v := rt.GlobalObject().Get(mockModuleGlobalKey)
	if v == nil || goja.IsNull(v) || goja.IsUndefined(v) {
		return h.HTTP.Request(ctx, method, url, args...)
	}

	orig, err := http.ToURL(url)
	if err != nil {
		return nil, err
	}

	vv, err := rt.RunString(mockModuleGlobalKey + ".resolve(\"" + orig.URL + "\")")
	if err != nil {
		return nil, err
	}

	rewrite := vv.ToString().String()
	if rewrite == orig.URL {
		return h.HTTP.Request(ctx, method, url, args...)
	}

	orig, err = httpext.NewURL(rewrite, orig.Name)
	if err != nil {
		return nil, err
	}

	return h.HTTP.Request(ctx, method, rt.ToValue(orig), args...)
}

func (h *HTTP) Get(ctx context.Context, url goja.Value, args ...goja.Value) (*http.Response, error) {
	args = append([]goja.Value{goja.Undefined()}, args...)

	return h.Request(ctx, http.HTTP_METHOD_GET, url, args...)
}

func (h *HTTP) Head(ctx context.Context, url goja.Value, args ...goja.Value) (*http.Response, error) {
	args = append([]goja.Value{goja.Undefined()}, args...)

	return h.Request(ctx, http.HTTP_METHOD_HEAD, url, args...)
}

func (h *HTTP) Post(ctx context.Context, url goja.Value, args ...goja.Value) (*http.Response, error) {
	return h.Request(ctx, http.HTTP_METHOD_POST, url, args...)
}

func (h *HTTP) Put(ctx context.Context, url goja.Value, args ...goja.Value) (*http.Response, error) {
	return h.Request(ctx, http.HTTP_METHOD_PUT, url, args...)
}

func (h *HTTP) Patch(ctx context.Context, url goja.Value, args ...goja.Value) (*http.Response, error) {
	return h.Request(ctx, http.HTTP_METHOD_PATCH, url, args...)
}

func (h *HTTP) Del(ctx context.Context, url goja.Value, args ...goja.Value) (*http.Response, error) {
	return h.Request(ctx, http.HTTP_METHOD_DELETE, url, args...)
}

func (h *HTTP) Options(ctx context.Context, url goja.Value, args ...goja.Value) (*http.Response, error) {
	return h.Request(ctx, http.HTTP_METHOD_OPTIONS, url, args...)
}
