package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConnections = 2
	defaultResponseTimeout    = 2 * time.Second
	defaultConnectionTimeout  = 2 * time.Second
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("content-type"), body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, errors.New("unable to create a request")
	}

	request.Header = fullHeaders

	client := c.getHttpClient()

	return client.Do(request)
}

func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	c.client = &http.Client{
		Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
			ResponseHeaderTimeout: c.getResponseTimeout(),
			DialContext: (&net.Dialer{
				Timeout: c.getConnectionTimeout(),
			}).DialContext,
		},
	}

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.requestTimeout > 0 {
		return c.requestTimeout
	}

	if c.disableTimeouts {
		return 0
	}

	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}

	if c.disableTimeouts {
		return 0
	}

	return defaultConnectionTimeout
}

func (c *httpClient) getRequestHeaders(requestHeader http.Header) http.Header {
	result := make(http.Header)

	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	for header, value := range requestHeader {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}
