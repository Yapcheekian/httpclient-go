package gohttp

import "net/http"

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeader http.Header) http.Header {
	result := make(http.Header)

	for header, value := range c.builder.headers {
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
