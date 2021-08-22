package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	requestTimeout     time.Duration
	disableTimeouts    bool
	client             *http.Client
}

type ClientBuilder interface {
	SetHeader(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(conn int) ClientBuilder
	SetRequestTimeout(timeout time.Duration) ClientBuilder
	DisableTimeouts(b bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder

	Build() Client
}

func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

func (c *clientBuilder) Build() Client {
	return &httpClient{
		builder: c,
	}
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetHeader(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(conn int) ClientBuilder {
	c.maxIdleConnections = conn
	return c
}

func (c *clientBuilder) SetRequestTimeout(timeout time.Duration) ClientBuilder {
	c.requestTimeout = timeout
	return c
}

func (c *clientBuilder) DisableTimeouts(b bool) ClientBuilder {
	c.disableTimeouts = b
	return c
}
