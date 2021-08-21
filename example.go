package main

import (
	"fmt"
	"net/http"

	"github.com/Yapcheekian/httpclient-go/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeader := make(http.Header)
	commonHeader.Set("Authorization", "Bearer ABC")
	client.SetHeader(commonHeader)

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
