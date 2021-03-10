package main

import (
	"fmt"

	"github.com/rwxrob/auth-go"
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("valid")
	x.Usage = `<name>`
	x.Summary = `print whether token is valid or not`

	x.Description = `
		The *valid* subcommand will print one of three values: 'true' if an
		access token exists for the given name and has not expired, 'false'
		if a token exists but has expired or has no refresh token, 'null' if
		there is no access token, and 'error' if an error prevents
		determining if it exists.`

	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}
		switch auth.Valid(args[0]) {
		case 1:
			fmt.Println("true")
		case 0:
			fmt.Println("false")
		case -1:
			fmt.Println("null")
		default:
			fmt.Println("error")
			return nil
		}
		return nil
	}
}
