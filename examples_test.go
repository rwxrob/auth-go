package auth_test

import (
	"fmt"

	"gitlab.com/rwxrob/auth"
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
