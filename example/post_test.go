package example

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/Yapcheekian/httpclient-go/gohttp"
)

func TestCreateRepo(t *testing.T) {
	t.Run("timeout from Github", func(t *testing.T) {
		gohttp.FlushMocks()

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		requestBody, _ := json.Marshal(repository)

		mock := gohttp.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: string(requestBody),

			Error: errors.New("timeout getting github endpoints"),
		}

		gohttp.AddMock(mock)

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
		gohttp.FlushMocks()

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		requestBody, _ := json.Marshal(repository)

		mock := gohttp.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: string(requestBody),

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id": 123, "name": "test-repo"}`,
		}

		gohttp.AddMock(mock)

		repo, err := CreateRepo(repository)

		if err != nil {
			t.Errorf("no error was expected but we got %s", err)
		}

		if repo == nil {
			t.Error("a repo was expected at this point")
		}
	})
}
