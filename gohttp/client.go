package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client             *http.Client
	Headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	requestTimeout     time.Duration
	disableTimeouts    bool
}

func New() HttpClient {
	return &httpClient{}
}

type HttpClient interface {
	SetHeader(headers http.Header)
	SetConnectionTimeout(timeout time.Duration)
	SetMaxIdleConnections(conn int)
	SetRequestTimeout(timeout time.Duration)
	DisableTimeouts(b bool)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeader(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

func (c *httpClient) SetMaxIdleConnections(conn int) {
	c.maxIdleConnections = conn
}

func (c *httpClient) SetRequestTimeout(timeout time.Duration) {
	c.requestTimeout = timeout
}

func (c *httpClient) DisableTimeouts(b bool) {
	c.disableTimeouts = b
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
