package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Yapcheekian/httpclient-go/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getGithubClient() gohttp.Client {
	builder := gohttp.NewBuilder()

	// builder.DisableTimeouts(true)
	commonHeader := make(http.Header)
	commonHeader.Set("Authorization", "Bearer ABC")

	client := builder.
		SetConnectionTimeout(2 * time.Second).
		SetMaxIdleConnections(2).
		SetRequestTimeout(5 * time.Second).
		SetHeader(commonHeader).
		Build()

	return client
}

func main() {
	getUrl()
	getUrl()
	getUrl()

	user := User{}
	createUser(user)
}

func getUrl() {
	resp, err := githubHttpClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
}

func createUser(user User) {
	resp, err := githubHttpClient.Post("https://api.github.com", nil, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
}
