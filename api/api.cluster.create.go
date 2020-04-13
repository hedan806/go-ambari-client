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

func newClusterCreateFunc(t Transport) ClusterCreate {
	return func(cluster string, body io.Reader, o ...func(*ClusterCreateRequest)) (*Response, error) {
		var r = ClusterCreateRequest{Cluster: cluster, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// ClusterCreate returns basic information about the cluster.
//
// See full documentation at http://www.elastic.co/guide/.
//
type ClusterCreate func(cluster string, body io.Reader, o ...func(*ClusterCreateRequest)) (*Response, error)

// ClusterCreateRequest configures the Cluster Create API request.
//
type ClusterCreateRequest struct {
	Cluster string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r ClusterCreateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("api") + 1 + len("v1") + 1 + len("clusters") + 1 + len(r.Cluster))
	path.WriteString("/")
	path.WriteString("api")
	path.WriteString("/")
	path.WriteString("v1")
	path.WriteString("/")
	path.WriteString("clusters")
	path.WriteString("/")
	path.WriteString(r.Cluster)

	params = make(map[string]string)

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
func (f ClusterCreate) WithContext(v context.Context) func(*ClusterCreateRequest) {
	return func(r *ClusterCreateRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f ClusterCreate) WithPretty() func(*ClusterCreateRequest) {
	return func(r *ClusterCreateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f ClusterCreate) WithHuman() func(*ClusterCreateRequest) {
	return func(r *ClusterCreateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f ClusterCreate) WithErrorTrace() func(*ClusterCreateRequest) {
	return func(r *ClusterCreateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f ClusterCreate) WithFilterPath(v ...string) func(*ClusterCreateRequest) {
	return func(r *ClusterCreateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f ClusterCreate) WithHeader(h map[string]string) func(*ClusterCreateRequest) {
	return func(r *ClusterCreateRequest) {
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
func (f ClusterCreate) WithOpaqueID(s string) func(*ClusterCreateRequest) {
	return func(r *ClusterCreateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
