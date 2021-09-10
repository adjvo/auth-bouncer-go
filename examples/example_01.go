package main

import (
	"fmt"
	"github.com/adjvo/auth-bouncer-go"
)

func main() {
	guard := authbouncer.NewGuard("http://local.api.auth.adjvo.com", "cAZNiqsS14ELm76Q0o1U2DvGzQcOrrmkyv6f7H5UkHodE9HSdwA9MhBfLZ5H", "Zyz3NxhWwRTFFVfKPYrU0nEqkXNY3KliJrG1i2eolq7xa5ETXrYcnQBToFyS")

	response := guard.Introspect("p3ibTOXjbXeVhRVW8VqJZmQ49E6GxMmDj5R5zXdAQNW6278pZRBgER0fqDIN")

	fmt.Println(response.Body.Data)
}
