package main

import (
	"fmt"

	"gitlab.com/rwxrob/auth"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("token")
	x.Usage = `<name>`
	x.Summary = `print the access (bearer) token, grant or refresh if needed`

	x.Description = `
		The *token* subcommand will print the access token (same as the
		*access* subcommand) but will also check to see if it needs to be
		refreshed or granted and if so calls those subcommands first before
		printing the token. This, therefore, involves potentially some
		blocking interactivity from the grant which should be taken into
		consideration before using this instead of *access* instead (which
		never triggers any interactivity, just returns whatever is cached).
		Another option to avoid any potential unwanted blocking
		interactivity is to call the *refresh* subcommand instead which
		always refreshes and prints the token and prints nothing if a token
		needs to be granted for any reason.`

	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}
		_, app, err := auth.Use(args[0])
		if err != nil {
			return err
		}
		fmt.Println(app.AccessToken)
		return nil
	}
}
