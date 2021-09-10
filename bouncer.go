package authbouncer

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/adjvo/auth-bouncer-go/authentication"
	"github.com/adjvo/auth-bouncer-go/introspection"
	"net/http"
)

type Guard struct {
	client   *http.Client
	token    *Token
	domain   string
	clientID string
	secret   string
}

// NewGuard instantiates new Guard struct
func NewGuard(domain, clientID, secret string) *Guard {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Guard{
		client:   &http.Client{Transport: tr},
		token:    NewToken(),
		domain:   domain,
		clientID: clientID,
		secret:   secret,
	}
}

// Introspect requests to Adjvo Auth server for token validation
func (g Guard) Introspect(token string) *introspection.Response {
	req := introspection.NewRequest("POST", g.domain+"/token/introspect")

	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", g.token.GetCachedToken()))

	req.Payload(introspection.Payload{token})

	resp, err := g.client.Do(req.Build())
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 401 {
		g.Authenticate()

		return g.Introspect(token)
	}

	return introspection.NewIntrospectionResponse(resp)
}

func (g Guard) Authenticate() {
	req := authentication.NewRequest("POST", g.domain+"/token")

	req.SetHeader("Content-Type", "application/json")

	req.Payload(authentication.Payload{
		GrantType:    "client_credentials",
		ClientID:     g.clientID,
		ClientSecret: g.secret,
	})

	resp, err := g.client.Do(req.Build())
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 401 {
		panic(errors.New(fmt.Sprintf("Client authentication failed with status code %d", resp.StatusCode)))
	}

	r := authentication.NewAuthenticationResponse(resp)

	g.token.SetCachedToken(r.Body.Data.AccessToken)
}
