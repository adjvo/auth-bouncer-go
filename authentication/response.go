package authentication

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Body struct {
	Data     Data
	Endpoint string
}

type Response struct {
	Body       Body
	StatusCode int
}

type Data struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	RequesterIP string `json:"requester_ip"`
	AccessToken string `json:"access_token"`
}

func NewAuthenticationResponse(resp *http.Response) *Response {
	var body Body

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(responseBody, &body); err != nil {
		panic(err)
	}

	return &Response{
		Body:       body,
		StatusCode: resp.StatusCode,
	}
}
