package introspection

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
	Active      bool      `json:"active"`
	AccessToken *string   `json:"access_token"`
	ClientID    *string   `json:"client_id"`
	UserID      *string   `json:"user_id"`
	Scopes      *[]string `json:"scopes"`
	User        *User     `json:"user"`
}

type User struct {
	Email       string `json:"email"`
	ActivatedAt string `json:"activated_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewIntrospectionResponse(resp *http.Response) *Response {
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
