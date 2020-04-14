// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 0.0.1: DO NOT EDIT

package api

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

func newClusterComponentsFunc(t Transport) ClusterComponents {
	return func(cluster string, o ...func(*ClusterComponentsRequest)) (*Response, error) {
		var r = ClusterComponentsRequest{Cluster: cluster}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// ClusterComponents returns basic information about the cluster.
//
// See full documentation at http://www.elastic.co/guide/.
//
type ClusterComponents func(cluster string, o ...func(*ClusterComponentsRequest)) (*Response, error)

// ClusterComponentsRequest configures the Cluster Components API request.
//
type ClusterComponentsRequest struct {
	Cluster string

	Fields          string
	MinimalResponse *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r ClusterComponentsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("api") + 1 + len("v1") + 1 + len("clusters") + 1 + len(r.Cluster) + 1 + len("components"))
	path.WriteString("/")
	path.WriteString("api")
	path.WriteString("/")
	path.WriteString("v1")
	path.WriteString("/")
	path.WriteString("clusters")
	path.WriteString("/")
	path.WriteString(r.Cluster)
	path.WriteString("/")
	path.WriteString("components")

	params = make(map[string]string)

	if r.Fields != "" {
		params["fields"] = r.Fields
	}

	if r.MinimalResponse != nil {
		params["minimal_response"] = strconv.FormatBool(*r.MinimalResponse)
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
func (f ClusterComponents) WithContext(v context.Context) func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
		r.ctx = v
	}
}

// WithFields - wait until the specified number of nodes is available.
//
func (f ClusterComponents) WithFields(v string) func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
		r.Fields = v
	}
}

// WithMinimalResponse - wait until the specified number of nodes is available.
//
func (f ClusterComponents) WithMinimalResponse(v bool) func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
		r.MinimalResponse = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f ClusterComponents) WithPretty() func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f ClusterComponents) WithHuman() func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f ClusterComponents) WithErrorTrace() func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f ClusterComponents) WithFilterPath(v ...string) func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f ClusterComponents) WithHeader(h map[string]string) func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
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
func (f ClusterComponents) WithOpaqueID(s string) func(*ClusterComponentsRequest) {
	return func(r *ClusterComponentsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
