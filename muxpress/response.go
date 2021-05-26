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
	"fmt"
	"net/http"
)

type Response struct {
	Writer http.ResponseWriter
}

func newResponse(w http.ResponseWriter) *Response {
	return &Response{Writer: w}
}

func (res *Response) Json(v interface{}) (*Response, error) { //nolint
	res.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	b, err := json.Marshal(v)
	if err != nil {
		return res, err
	}

	_, err = res.Writer.Write(b)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (res *Response) Text(format string, v ...interface{}) (*Response, error) {
	res.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")

	_, err := res.Writer.Write([]byte(fmt.Sprintf(format, v...)))
	if err != nil {
		return res, err
	}

	return res, nil
}

func (res *Response) Html(b []byte) (*Response, error) { //nolint
	res.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err := res.Writer.Write(b); err != nil {
		return res, err
	}

	return res, nil
}

func (res *Response) Binary(b []byte) (*Response, error) {
	res.Writer.Header().Set("Content-Type", "application/octet-stream")

	if _, err := res.Writer.Write(b); err != nil {
		return res, err
	}

	return res, nil
}

func (res *Response) Send(v interface{}) (*Response, error) {
	switch val := v.(type) {
	case string:
		return res.Html([]byte(val))
	case []byte:
		return res.Binary(val)
	default:
		return res.Json(v)
	}
}

func (res *Response) Status(code int) *Response {
	res.Writer.WriteHeader(code)

	return res
}

func (res *Response) Type(mime string) *Response {
	res.Writer.Header().Set("Content-Type", mime)

	return res
}

func (res *Response) Vary(header string) *Response {
	res.Writer.Header().Add("Vary", header)

	return res
}

func (res *Response) Set(field string, value string) *Response {
	res.Writer.Header().Set(field, value)

	return res
}

func (res *Response) Append(field string, value string) *Response {
	res.Writer.Header().Add(field, value)

	return res
}

func (res *Response) Redirect(code int, loc string) *Response {
	res.Writer.WriteHeader(code)
	res.Writer.Header().Set("Location", loc)

	return res
}

/*
type cookieOptions struct {
	Domain   string
	Expires  int
	HttpOnly bool
	MaxAge   int
	Path     string
	Secure   bool
}

func (res *response) Cookie(name string, v interface{}, opts *cookieOptions) {
}

*/
