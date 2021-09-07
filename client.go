package authbouncer

type Guard struct {
	Domain   string
	ClientID string
	Secret   string
}

func (Guard) Introspect(token string) {}