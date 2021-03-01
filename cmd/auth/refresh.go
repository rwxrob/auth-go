package main

import (
	"fmt"
	"time"

	"github.com/rwxrob/auth"
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("refresh")
	x.Usage = `<name>`
	x.Summary = `refresh the access token for given entry`

	x.Description = `
		The *refresh* subcommand forces an immediate refresh of the access
		token using the provide refresh token if the expiry has passed and
		prints it to standard output. This is useful when finer control is
		required of the authorization process and when blocking
		interactivity of a full *token* subcommand is not wanted.`

	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		config, app, err := auth.Lookup(args[0])
		if err != nil {
			return err
		}
		if app.RefreshToken != "" && !app.Expiry.Before(time.Now()) {
			fmt.Println(app.AccessToken)
			return nil
		}
		err = app.Refresh()
		if err != nil {
			return err
		}
		fmt.Println(app.AccessToken)
		return config.Store()
	}
}
