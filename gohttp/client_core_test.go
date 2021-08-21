package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeader(t *testing.T) {
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("content-type", "application/json")
	commonHeaders.Set("user-agent", "test-agent")
	client.builder.SetHeader(commonHeaders)

	requestHeaders := make(http.Header)
	requestHeaders.Set("x-request-id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers")
	}

	if finalHeaders.Get("x-request-id") != "ABC-123" {
		t.Error("we expect x-request-id to be ABC-123")
	}

	if finalHeaders.Get("user-agent") != "test-agent" {
		t.Error("we expect user-agent to be test-agent")
	}

	if finalHeaders.Get("content-type") != "application/json" {
		t.Error("we expect content-type to be application/json")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}
	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if body != nil {
			t.Error("no body expected when passing nil body")
		}
	})

	t.Run("BodyWithJSON", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("no error expected when marshalling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json obtained")
		}
	})

	t.Run("BodyWithXML", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/xml", requestBody)

		if err != nil {
			t.Error("no error expected when marshalling slice as json")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("invalid json obtained")
		}
	})

	t.Run("BodyWithDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("", requestBody)

		if err != nil {
			t.Error("no error expected when marshalling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json obtained")
		}
	})
}
