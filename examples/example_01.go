package main

import (
	"fmt"
	"github.com/adjvo/auth-bouncer-go"
)

// Auth example app handler
type Auth struct {
	authbouncer.Guard
}

func main() {
	auth := Auth{
		Guard: authbouncer.NewGuard("http://local.api.auth.adjvo.com", "cAZNiqsS14ELm76Q0o1U2DvGzQcOrrmkyv6f7H5UkHodE9HSdwA9MhBfLZ5H", "Zyz3NxhWwRTFFVfKPYrU0nEqkXNY3KliJrG1i2eolq7xa5ETXrYcnQBToFyS"),
	}

	i, err := auth.Introspect("jdQkQGu0XY6ygmlwltFnpix4fXYZXtco6iGAel4BtsO1jCWZyJ6V2ZBW5ihe")
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
