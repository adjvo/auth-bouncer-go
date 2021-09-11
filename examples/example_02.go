package main

import (
	"fmt"
	"github.com/adjvo/auth-bouncer-go"
)

// MockTesting example app handler
type MockTesting struct {
	authbouncer.Guard
}

func main() {
	auth := MockTesting{
		Guard: authbouncer.Mock{
			AccessToken: "TEST",
			ClientID:    "TEST",
			UserID:      "TEST",
			Scopes:      []string{"TEST"},
		},
	}

	introspection, err := auth.Introspect("p3ibTOXjbXeVhRVW8VqJZmQ49E6GxMmDj5R5zXdAQNW6278pZRBgER0fqDIN")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(introspection.AccessToken)
	fmt.Println(introspection.ClientID)
	fmt.Println(introspection.UserID)
	fmt.Println(introspection.Scopes)
}
