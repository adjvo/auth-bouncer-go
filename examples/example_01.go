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

	introspection, err := auth.Introspect("hHR0F8OrhlSoXwVUau91Dl2onEMS8zTOZnhxyzRJ51B0vLkbKLNEq1ejGicb")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(introspection.AccessToken)
	fmt.Println(introspection.ClientID)
	fmt.Println(introspection.UserID)
	fmt.Println(introspection.Scopes)
}
