// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 0.0.1: DO NOT EDIT

package api

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newBlueprintCreateFunc(t Transport) BlueprintCreate {
	return func(name string, body io.Reader, o ...func(*BlueprintCreateRequest)) (*Response, error) {
		var r = BlueprintCreateRequest{Name: name, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// BlueprintCreate returns basic information about the cluster.
//
// See full documentation at http://www.elastic.co/guide/.
//
type BlueprintCreate func(name string, body io.Reader, o ...func(*BlueprintCreateRequest)) (*Response, error)

// BlueprintCreateRequest configures the Blueprint Create API request.
//
type BlueprintCreateRequest struct {
	Body io.Reader

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
func (r BlueprintCreateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

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

	req, err := newRequest(method, path.String(), r.Body)
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

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
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
func (f BlueprintCreate) WithContext(v context.Context) func(*BlueprintCreateRequest) {
	return func(r *BlueprintCreateRequest) {
		r.ctx = v
	}
}

// WithFields - wait until the specified number of nodes is available.
//
func (f BlueprintCreate) WithFields(v string) func(*BlueprintCreateRequest) {
	return func(r *BlueprintCreateRequest) {
		r.Fields = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f BlueprintCreate) WithPretty() func(*BlueprintCreateRequest) {
	return func(r *BlueprintCreateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f BlueprintCreate) WithHuman() func(*BlueprintCreateRequest) {
	return func(r *BlueprintCreateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f BlueprintCreate) WithErrorTrace() func(*BlueprintCreateRequest) {
	return func(r *BlueprintCreateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f BlueprintCreate) WithFilterPath(v ...string) func(*BlueprintCreateRequest) {
	return func(r *BlueprintCreateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f BlueprintCreate) WithHeader(h map[string]string) func(*BlueprintCreateRequest) {
	return func(r *BlueprintCreateRequest) {
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
func (f BlueprintCreate) WithOpaqueID(s string) func(*BlueprintCreateRequest) {
	return func(r *BlueprintCreateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
