package ambari

import (
	"encoding/base64"
	"fmt"
	"github.com/hedan806/go-ambari-client/api"
	"github.com/hedan806/go-ambari-client/transport"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	defaultURL = "http://localhost:8080"
)

type Config struct {
	Address  string
	Username string
	Password string

	EnableMetrics     bool // Enable the metrics collection.
	EnableDebugLogger bool // Enable the debug logging.

	Logger transport.Logger

	Transport http.RoundTripper  // The HTTP transport object.
	Selector  transport.Selector // The selector object.

	// Optional constructor function for a custom ConnectionPool. Default: nil.
	ConnectionPoolFunc func([]*transport.Connection, transport.Selector) transport.ConnectionPool
}

type Client struct {
	*api.API
	Transport transport.Interface
}

func NewDefaultClient() (*Client, error) {
	return NewClient(Config{})
}

func NewClient(cfg Config) (*Client, error) {
	var addrs []string

	if len(cfg.Address) > 0 {
		addrs = append(addrs, cfg.Address)
	}

	urls, err := addrsToURLs(addrs)
	if err != nil {
		return nil, fmt.Errorf("cannot create client: %s", err)
	}

	if len(urls) == 0 {
		u, _ := url.Parse(defaultURL) // errcheck exclude
		urls = append(urls, u)
	}

	// TODO(karmi): Refactor
	if urls[0].User != nil {
		cfg.Username = urls[0].User.Username()
		pw, _ := urls[0].User.Password()
		cfg.Password = pw
	}

	tp, err := transport.New(transport.Config{
		URLs:     urls,
		Username: cfg.Username,
		Password: cfg.Password,

		EnableDebugLogger: cfg.EnableDebugLogger,
		Logger:            cfg.Logger,

		Transport:          cfg.Transport,
		Selector:           cfg.Selector,
		ConnectionPoolFunc: cfg.ConnectionPoolFunc,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating transport: %s", err)
	}

	client := &Client{Transport: tp, API: api.New(tp)}

	return client, nil
}

// Perform delegates to Transport to execute a request and return a response.
//
func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	return c.Transport.Perform(req)
}

// addrsFromEnvironment returns a list of addresses by splitting
// the ELASTICSEARCH_URL environment variable with comma, or an empty list.
//
func addrsFromEnvironment() []string {
	var addrs []string

	if envURLs, ok := os.LookupEnv("ELASTICSEARCH_URL"); ok && envURLs != "" {
		list := strings.Split(envURLs, ",")
		for _, u := range list {
			addrs = append(addrs, strings.TrimSpace(u))
		}
	}

	return addrs
}

// addrsToURLs creates a list of url.URL structures from url list.
//
func addrsToURLs(addrs []string) ([]*url.URL, error) {
	var urls []*url.URL
	for _, addr := range addrs {
		u, err := url.Parse(strings.TrimRight(addr, "/"))
		if err != nil {
			return nil, fmt.Errorf("cannot parse url: %v", err)
		}

		urls = append(urls, u)
	}
	return urls, nil
}

// addrFromCloudID extracts the Elasticsearch URL from CloudID.
// See: https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html
//
func addrFromCloudID(input string) (string, error) {
	var scheme = "https://"

	values := strings.Split(input, ":")
	if len(values) != 2 {
		return "", fmt.Errorf("unexpected format: %q", input)
	}
	data, err := base64.StdEncoding.DecodeString(values[1])
	if err != nil {
		return "", err
	}
	parts := strings.Split(string(data), "$")

	if len(parts) < 2 {
		return "", fmt.Errorf("invalid encoded value: %s", parts)
	}

	return fmt.Sprintf("%s%s.%s", scheme, parts[1], parts[0]), nil
}
