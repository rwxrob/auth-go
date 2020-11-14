package auth_test

import (
	"fmt"
	"os"

	"gitlab.com/rwxrob/auth-go"
)

func ExampleGet() {
	os.Setenv("AUTHDIR", "testdata")
	id := auth.Get("some.io/client/id")
	sec := auth.Get("some.io/client/secret")
	fmt.Println(id)
	fmt.Println(sec)
	// Output:
	// some_client_id
	// some_client_secret
}
