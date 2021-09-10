package introspection

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Payload struct {
	Token string `json:"token"`
}

type Request struct {
	Method  string
	Url     string
	Headers map[string]string
	payload io.Reader
}

func NewRequest(requestMethod, url string) Request {
	return Request{
		Method:  requestMethod,
		Url:     url,
		Headers: make(map[string]string),
	}
}

func (r *Request) Payload(p interface{}) {
	jsonValue, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	r.payload = bytes.NewBuffer(jsonValue)
}

func (r *Request) SetHeader(key, value string) {
	r.Headers[key] = value
}

func (r *Request) Build() *http.Request {
	req, err := http.NewRequest(r.Method, r.Url, r.payload)
	if err != nil {
		panic(err)
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	return req
}
