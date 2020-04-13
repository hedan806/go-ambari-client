package transport

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	userAgent   string
	reGoVersion = regexp.MustCompile(`go(\d+\.\d+\..+)`)

	defaultMaxRetries    = 3
	defaultRetryOnStatus = [...]int{502, 503, 504}
)

// Interface defines the interface for HTTP client.
//
type Interface interface {
	Perform(*http.Request) (*http.Response, error)
}

// Config represents the configuration of HTTP client.
//
type Config struct {
	URLs     []*url.URL
	Username string
	Password string

	EnableDebugLogger bool

	Transport http.RoundTripper
	Logger    Logger
	Selector  Selector

	ConnectionPoolFunc func([]*Connection, Selector) ConnectionPool
}

// Client represents the HTTP client.
//
type Client struct {
	sync.Mutex

	urls     []*url.URL
	username string
	password string

	transport http.RoundTripper
	logger    Logger
	selector  Selector
	pool      ConnectionPool
	poolFunc  func([]*Connection, Selector) ConnectionPool
}

// New creates new transport client.
//
// http.DefaultTransport will be used if no transport is passed in the configuration.
//
func New(cfg Config) (*Client, error) {
	if cfg.Transport == nil {
		cfg.Transport = http.DefaultTransport
	}

	var conns []*Connection
	for _, u := range cfg.URLs {
		conns = append(conns, &Connection{URL: u})
	}

	client := Client{
		urls:     cfg.URLs,
		username: cfg.Username,
		password: cfg.Password,

		logger: cfg.Logger,

		transport: cfg.Transport,
		selector:  cfg.Selector,
		poolFunc:  cfg.ConnectionPoolFunc,
	}

	if client.poolFunc != nil {
		client.pool = client.poolFunc(conns, client.selector)
	} else {
		client.pool, _ = NewConnectionPool(conns, client.selector)
	}

	return &client, nil
}

// Perform executes the request and returns a response or error.
//
func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	var (
		res *http.Response
		err error
	)

	// Update request
	c.setReqUserAgent(req)

	if req.Body != nil && req.Body != http.NoBody && req.GetBody == nil {
	}

	//for i := 1; i <= c.maxRetries; i++ {
	for i := 1; i <= 1; i++ {
		var (
			conn        *Connection
			shouldRetry bool
		)

		// Get connection from the pool
		c.Lock()
		conn, err = c.pool.Next()
		c.Unlock()
		if err != nil {
			return nil, fmt.Errorf("cannot get connection: %s", err)
		}

		// Update request
		c.setReqURL(conn.URL, req)
		c.setReqAuth(conn.URL, req)

		// Set up time measures and execute the request
		start := time.Now().UTC()
		res, err = c.transport.RoundTrip(req)
		dur := time.Since(start)

		// Log request and response
		if c.logger != nil {
			if c.logger.RequestBodyEnabled() && req.Body != nil && req.Body != http.NoBody {
				req.Body, _ = req.GetBody()
			}
			c.logRoundTrip(req, res, err, start, dur)
		}

		if err != nil {

			// Report the connection as unsuccessful
			c.Lock()
			c.pool.OnFailure(conn)
			c.Unlock()

			// Retry on EOF errors
			if err == io.EOF {
				shouldRetry = true
			}

		} else {
			// Report the connection as succesfull
			c.Lock()
			c.pool.OnSuccess(conn)
			c.Unlock()
		}

		// Break if retry should not be performed
		if !shouldRetry {
			break
		}

	}

	// TODO(karmi): Wrap error
	return res, err
}

// URLs returns a list of transport URLs.
//
//
func (c *Client) URLs() []*url.URL {
	return c.pool.URLs()
}

func (c *Client) setReqURL(u *url.URL, req *http.Request) *http.Request {
	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host

	if u.Path != "" {
		var b strings.Builder
		b.Grow(len(u.Path) + len(req.URL.Path))
		b.WriteString(u.Path)
		b.WriteString(req.URL.Path)
		req.URL.Path = b.String()
	}

	return req
}

func (c *Client) setReqAuth(u *url.URL, req *http.Request) *http.Request {
	if _, ok := req.Header["Authorization"]; !ok {
		if u.User != nil {
			password, _ := u.User.Password()
			req.SetBasicAuth(u.User.Username(), password)
			return req
		}

		if c.username != "" && c.password != "" {
			req.SetBasicAuth(c.username, c.password)
			return req
		}
	}

	return req
}

func (c *Client) setReqUserAgent(req *http.Request) *http.Request {
	req.Header.Set("User-Agent", userAgent)
	return req
}

func (c *Client) logRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) {
	var dupRes http.Response
	if res != nil {
		dupRes = *res
	}
	if c.logger.ResponseBodyEnabled() {
		if res != nil && res.Body != nil && res.Body != http.NoBody {
			b1, b2, _ := duplicateBody(res.Body)
			dupRes.Body = b1
			res.Body = b2
		}
	}
	c.logger.LogRoundTrip(req, &dupRes, err, start, dur) // errcheck exclude
}

func initUserAgent() string {
	var b strings.Builder

	b.WriteString("go-elasticsearch")
	b.WriteRune('/')
	b.WriteString("Version")
	b.WriteRune(' ')
	b.WriteRune('(')
	b.WriteString(runtime.GOOS)
	b.WriteRune(' ')
	b.WriteString(runtime.GOARCH)
	b.WriteString("; ")
	b.WriteString("Go ")
	if v := reGoVersion.ReplaceAllString(runtime.Version(), "$1"); v != "" {
		b.WriteString(v)
	} else {
		b.WriteString(runtime.Version())
	}
	b.WriteRune(')')

	return b.String()
}
