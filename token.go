package authbouncer

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Token struct {
	*cache.Cache
}

func NewToken() *Token {
	c := cache.New(5*time.Minute, 10*time.Minute)

	return &Token{
		c,
	}
}

// GetCachedToken fetches client access token from cache
func (t Token) GetCachedToken() string {
	if token, found := t.Get("token"); found {
		return token.(string)
	}

	return ""
}

// SetCachedToken stores client access token to cache
func (t Token) SetCachedToken(token string) {
	t.Set("token", token, cache.DefaultExpiration)
}
