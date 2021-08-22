package gohttpmock

import (
	"fmt"
	"net/http"

	"github.com/Yapcheekian/httpclient-go/core"
)

// Mock struct configure HTTP mocks based on the combination
// between request method, request url and request body.
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	ResponseBody       string
	ResponseStatusCode int
	Error              error
}

// GetResponse returns an response object based on the mock configuration.
func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	response := core.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
	}

	return &response, nil
}
