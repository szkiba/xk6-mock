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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

type dict map[string]interface{}

type Request struct {
	Request  *http.Request
	Params   map[string]string
	Cookies  map[string]string
	Path     string
	Protocol string
	Method   string
	Query    dict
	Body     dict
}

func newRequest(r *http.Request) (req *Request, err error) {
	req = new(Request)

	req.Request = r
	req.Cookies = mapCookies(r.Cookies())

	req.Params = mux.Vars(r)

	if req.Params == nil {
		req.Params = map[string]string{}
	}

	req.Method = r.Method
	req.Path = r.URL.Path
	req.Protocol = r.URL.Scheme
	req.Query = mapValues(r.URL.Query())

	if req.Body, err = parseBody(r); err != nil {
		return nil, err
	}

	return req, nil
}

func (req *Request) Get(field string) string {
	return req.Request.Header.Get(field)
}

func (req *Request) Header(field string) string {
	return req.Request.Header.Get(field)
}

func mapCookies(all []*http.Cookie) map[string]string {
	cookies := make(map[string]string, len(all))

	for _, c := range all {
		cookies[c.Name] = c.Value
	}

	return cookies
}

func mapValues(in url.Values) dict {
	out := make(dict)

	if in == nil {
		return out
	}

	for k, v := range in {
		if len(v) == 1 {
			out[k] = v[0]
		} else {
			out[k] = v
		}
	}

	return out
}

func parseBody(r *http.Request) (dict, error) {
	if r.ContentLength == 0 || !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		return nil, nil
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	out := dict{}

	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}

	return out, nil
}
