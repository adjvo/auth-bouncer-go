package main

import (
	"fmt"
	"github.com/adjvo/auth-bouncer-go"
	"github.com/adjvo/auth-bouncer-go/introspection"
)

// MockTesting example app handler
type MockTesting struct {
	authbouncer.Guard
}

func main() {
	auth := MockTesting{
		Guard: authbouncer.Mock{
			AccessToken: "access-token",
			ClientID:    "client-id",
			UserID:      "user-id",
			Scopes:      []string{"scope"},
			User: introspection.User{
				Email:       "test@email.com",
				ActivatedAt: "2021-05-30 10:59:45",
				CreatedAt:   "2021-05-30 10:59:45",
				UpdatedAt:   "2021-05-30 10:59:45",
			},
		},
	}

	i, err := auth.Introspect("p3ibTOXjbXeVhRVW8VqJZmQ49E6GxMmDj5R5zXdAQNW6278pZRBgER0fqDIN")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(i.AccessToken)
	fmt.Println(i.ClientID)
	fmt.Println(i.UserID)
	fmt.Println(i.Scopes)
	fmt.Println(i.User)
}
