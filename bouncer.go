package authbouncer

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/adjvo/auth-bouncer-go/authentication"
	"github.com/adjvo/auth-bouncer-go/introspection"
	"net/http"
)

type Guard interface {
	Introspect(token string) (Introspect, error)
}

type guard struct {
	client   *http.Client
	token    *Token
	domain   string
	clientID string
	secret   string
}

type Introspect struct {
	AccessToken string   `json:"access_token"`
	ClientID    string   `json:"client_id"`
	UserID      string   `json:"user_id"`
	Scopes      []string `json:"scopes"`
}

// NewGuard instantiates new guard struct
func NewGuard(domain, clientID, secret string) *guard {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &guard{
		client:   &http.Client{Transport: tr},
		token:    NewToken(),
		domain:   domain,
		clientID: clientID,
		secret:   secret,
	}
}

// Introspect requests to Adjvo Auth server for token validation
func (g guard) Introspect(token string) (Introspect, error) {
	var introspect Introspect

	req := introspection.NewRequest("POST", g.domain+"/token/introspect")

	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", g.token.GetCachedToken()))

	req.Payload(introspection.Payload{token})

	resp, err := g.client.Do(req.Build())
	if err != nil {
		return introspect, err
	}

	if resp.StatusCode == 401 {
		err := g.Authenticate()
		if err != nil {
			return introspect, err
		}

		return g.Introspect(token)
	}

	r := introspection.NewIntrospectionResponse(resp)

	if !r.Body.Data.Active {
		return introspect, errors.New("token expired or invalid")
	}

	introspect.AccessToken = *r.Body.Data.AccessToken
	introspect.ClientID = *r.Body.Data.ClientID
	introspect.UserID = *r.Body.Data.UserID
	introspect.Scopes = *r.Body.Data.Scopes

	return introspect, nil
}

// Authenticate  requests to Adjvo Auth server for client access token
func (g guard) Authenticate() error {
	req := authentication.NewRequest("POST", g.domain+"/token")

	req.SetHeader("Content-Type", "application/json")

	req.Payload(authentication.Payload{
		GrantType:    "client_credentials",
		ClientID:     g.clientID,
		ClientSecret: g.secret,
	})

	resp, err := g.client.Do(req.Build())
	if err != nil {
		return err
	}

	if resp.StatusCode == 401 {
		return errors.New("invalid client credentials")
	}

	r := authentication.NewAuthenticationResponse(resp)

	g.token.SetCachedToken(r.Body.Data.AccessToken)

	return nil
}
