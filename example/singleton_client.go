package example

import (
	"time"

	"github.com/Yapcheekian/httpclient-go/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetMaxIdleConnections(5).
		SetRequestTimeout(5 * time.Second).
		Build()

	return client
}
