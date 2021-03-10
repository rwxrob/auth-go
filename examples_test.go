package auth_test

import (
	"fmt"

	"github.com/rwxrob/auth-go"
)

func ExampleLookup() {
	config, app, err := auth.Lookup("some")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Has("some"))
	fmt.Println(err)
	fmt.Println(app.Name)
}
