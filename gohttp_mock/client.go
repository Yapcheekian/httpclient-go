package gohttpmock

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct{}

func (c *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	requestBody, err := req.GetBody()

	if err != nil {
		return nil, err
	}

	defer requestBody.Close()

	body, err := ioutil.ReadAll(requestBody)

	if err != nil {
		return nil, err
	}

	var response http.Response

	mock := mockUpServer.mocks[mockUpServer.getMockKey(req.Method, req.URL.String(), string(body))]

	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}

		response.StatusCode = mock.ResponseStatusCode
		response.Body = ioutil.NopCloser(strings.NewReader(mock.ResponseBody))
		response.ContentLength = int64(len(mock.ResponseBody))
		response.Request = req

		return &response, nil
	}

	return nil, errors.New("no mock matching")
}
