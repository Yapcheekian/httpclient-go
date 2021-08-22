package example

import "fmt"

/*
{
	"current_user_url": "https://api.github.com/user",
	"current_user_authorizations_html_url": "https://github.com/settings/connections/applications{/client_id}",
	"authorizations_url": "https://api.github.com/authorizations"
}
*/

type Endpoints struct {
	CurrentUserUrl                   string `json:"current_user_url"`
	CurrentUserAuthorizationsHTMLUrl string `json:"current_user_authorizations_html_url"`
	AuthorizationsUrl                string `json:"authorizations_url"`
}

func GetEndpoints() (*Endpoints, error) {
	resp, err := httpClient.Get("https://api.github.com", nil)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Status code: %d\n", resp.StatusCode())
	fmt.Printf("Status: %s\n", resp.Status())
	fmt.Printf("Header: %v\n", resp.Headers())
	fmt.Printf("Body: %v\n", resp.String())

	var endpoints Endpoints

	if err := resp.UnmarshalJson(&endpoints); err != nil {
		panic(err)
	}

	fmt.Println(endpoints.CurrentUserUrl)
	fmt.Println(endpoints.AuthorizationsUrl)

	return &endpoints, nil
}
