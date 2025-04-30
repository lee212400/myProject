package saas

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var client *http.Client

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getClient() *http.Client {
	if client == nil {
		client = &http.Client{
			Timeout:   10 * time.Second,
			Transport: getTransport(),
		}
	}

	return client
}

func getTransport() http.RoundTripper {
	return &http.Transport{
		TLSHandshakeTimeout: 5 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		IdleConnTimeout:       90 * time.Second,
		MaxIdleConns:          100,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func GetUser(id string) (string, int, error) {

	uri := "https://sample/users/" + id

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := getClient().Do(req)
	if err != nil {
		return "", 0, err
	} else if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("error status code:%d", resp.StatusCode)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	fmt.Println(string(bodyBytes))

	dt := &User{}

	if err := json.Unmarshal(bodyBytes, &dt); err != nil {
		return "", 0, err
	}

	return dt.Name, dt.Age, nil
}
