package authbouncer

type Mock struct {
	AccessToken string
	ClientID    string
	UserID      string
	Scopes      []string
}

// Introspect mocks the bouncer's Introspect function
func (m Mock) Introspect(string) (Introspect, error) {
	introspect := Introspect{
		AccessToken: m.AccessToken,
		ClientID:    m.ClientID,
		UserID:      m.UserID,
		Scopes:      m.Scopes,
	}

	return introspect, nil
}
