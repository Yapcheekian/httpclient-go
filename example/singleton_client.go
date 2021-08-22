package example

import (
	"net/http"
	"time"

	"github.com/Yapcheekian/httpclient-go/gohttp"
	"github.com/Yapcheekian/httpclient-go/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetMaxIdleConnections(5).
		SetRequestTimeout(5 * time.Second).
		SetHeader(headers).
		Build()

	return client
}
