package authbouncer

import "github.com/adjvo/auth-bouncer-go/introspection"

type Mock struct {
	AccessToken string
	ClientID    string
	UserID      string
	Scopes      []string
	User        introspection.User
}

// Introspect mocks the bouncer's Introspect function
func (m Mock) Introspect(string) (Introspect, error) {
	introspect := Introspect{
		AccessToken: m.AccessToken,
		ClientID:    m.ClientID,
		UserID:      m.UserID,
		Scopes:      m.Scopes,
		User:        m.User,
	}

	return introspect, nil
}
