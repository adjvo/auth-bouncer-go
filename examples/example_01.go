package main

import (
	"fmt"
	"github.com/adjvo/auth-bouncer-go"
)

func main() {
	guard := authbouncer.NewGuard("http://local.api.auth.adjvo.com", "cAZNiqsS14ELm76Q0o1U2DvGzQcOrrmkyv6f7H5UkHodE9HSdwA9MhBfLZ5H", "Zyz3NxhWwRTFFVfKPYrU0nEqkXNY3KliJrG1i2eolq7xa5ETXrYcnQBToFyS")

	introspection, err := guard.Introspect("p3ibTOXjbXeVhRVW8VqJZmQ49E6GxMmDj5R5zXdAQNW6278pZRBgER0fqDIN")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(introspection.AccessToken)
	fmt.Println(introspection.ClientID)
	fmt.Println(introspection.UserID)
	fmt.Println(introspection.Scopes)
}
