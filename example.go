package httpclient_go

import (
	"fmt"

	"github.com/Yapcheekian/httpclient-go/gohttp"
)

func basicExample() {
	client := gohttp.New()

	resp, err := client.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
}
