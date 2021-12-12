package authentication

import (
	"encoding/json"
	"io"
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

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
