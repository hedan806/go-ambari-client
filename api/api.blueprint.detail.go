// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 0.0.1: DO NOT EDIT

package api

import (
	"context"
	"net/http"
	"strings"
)

func newBlueprintDetailFunc(t Transport) BlueprintDetail {
	return func(name string, o ...func(*BlueprintDetailRequest)) (*Response, error) {
		var r = BlueprintDetailRequest{Name: name}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// BlueprintDetail returns basic information about the cluster.
//
// See full documentation at http://www.elastic.co/guide/.
//
type BlueprintDetail func(name string, o ...func(*BlueprintDetailRequest)) (*Response, error)

// BlueprintDetailRequest configures the Blueprint Detail API request.
//
type BlueprintDetailRequest struct {
	Name string

	Fields string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r BlueprintDetailRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("api") + 1 + len("v1") + 1 + len("blueprints") + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("api")
	path.WriteString("/")
	path.WriteString("v1")
	path.WriteString("/")
	path.WriteString("blueprints")
	path.WriteString("/")
	path.WriteString(r.Name)

	params = make(map[string]string)

	if r.Fields != "" {
		params["fields"] = r.Fields
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f BlueprintDetail) WithContext(v context.Context) func(*BlueprintDetailRequest) {
	return func(r *BlueprintDetailRequest) {
		r.ctx = v
	}
}

// WithFields - wait until the specified number of nodes is available.
//
func (f BlueprintDetail) WithFields(v string) func(*BlueprintDetailRequest) {
	return func(r *BlueprintDetailRequest) {
		r.Fields = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f BlueprintDetail) WithPretty() func(*BlueprintDetailRequest) {
	return func(r *BlueprintDetailRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f BlueprintDetail) WithHuman() func(*BlueprintDetailRequest) {
	return func(r *BlueprintDetailRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f BlueprintDetail) WithErrorTrace() func(*BlueprintDetailRequest) {
	return func(r *BlueprintDetailRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f BlueprintDetail) WithFilterPath(v ...string) func(*BlueprintDetailRequest) {
	return func(r *BlueprintDetailRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f BlueprintDetail) WithHeader(h map[string]string) func(*BlueprintDetailRequest) {
	return func(r *BlueprintDetailRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f BlueprintDetail) WithOpaqueID(s string) func(*BlueprintDetailRequest) {
	return func(r *BlueprintDetailRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
