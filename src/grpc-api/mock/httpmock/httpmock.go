package httpmock

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

const (
	Response200 int = iota
	Response400
	ResponseNoBody
)

func GetUser(userId string, testName string, testAge int, tCase int) {
	uri := "https://sample/users/" + userId

	if tCase == Response200 {
		httpmock.RegisterResponder("GET", uri,
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, map[string]any{
					"name": testName, "age": testAge,
				})
				resp.Header.Set("Content-Type", "application/json")
				return resp, err
			})
	} else if tCase == Response400 {
		httpmock.RegisterResponder("GET", uri,
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(400, map[string]any{})
				resp.Header.Set("Content-Type", "application/json")
				return resp, err
			})
	} else if tCase == ResponseNoBody {
		httpmock.RegisterResponder("GET", uri,
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(500, nil)
				resp.Header.Set("Content-Type", "application/json")
				return nil, err
			})
	}

}
