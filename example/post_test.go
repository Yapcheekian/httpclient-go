package example

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	gohttpmock "github.com/Yapcheekian/httpclient-go/gohttp_mock"
)

func TestCreateRepo(t *testing.T) {
	t.Run("timeout from Github", func(t *testing.T) {
		gohttpmock.DeleteMocks()

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		requestBody, _ := json.Marshal(repository)

		mock := gohttpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: string(requestBody),

			Error: errors.New("timeout getting github endpoints"),
		}

		gohttpmock.AddMock(mock)

		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("no repo expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message")
		}
	})

	t.Run("no error", func(t *testing.T) {
		gohttpmock.DeleteMocks()

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		requestBody, _ := json.Marshal(repository)

		mock := gohttpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: string(requestBody),

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id": 123, "name": "test-repo"}`,
		}

		gohttpmock.AddMock(mock)

		repo, err := CreateRepo(repository)

		if err != nil {
			t.Errorf("no error was expected but we got %s", err)
		}

		if repo == nil {
			t.Error("a repo was expected at this point")
		}
	})
}
