package example

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"testing"

	gohttpmock "github.com/Yapcheekian/httpclient-go/gohttp_mock"
)

func TestMain(m *testing.M) {
	gohttpmock.StartMockServer()
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		mock := gohttpmock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		}
		gohttpmock.DeleteMocks()
		gohttpmock.AddMock(mock)

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		mock := gohttpmock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		}

		gohttpmock.DeleteMocks()
		gohttpmock.AddMock(mock)

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field Endpoints.current_user_url of type string") {
			t.Error("invalid error message")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		mock := gohttpmock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		}

		gohttpmock.DeleteMocks()
		gohttpmock.AddMock(mock)

		endpoints, err := GetEndpoints()

		if err != nil {
			t.Errorf("no error was expected but we got %s\n", err)
		}

		if endpoints == nil {
			t.Error("endpoints were expected but we got nil")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})
}
