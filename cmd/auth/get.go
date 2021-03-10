package main

import (
	"fmt"
	"os"

	"github.com/rwxrob/auth-go"
	"github.com/rwxrob/cmdtab"
)

func init() {
	// FIXME cmdtab support for subcommands of subcommands is bork
	/*
		x := cmdtab.New("get", "conf", "access", "refresh", "type",
			"expiry", "state", "code", "id", "secret", "scopes", "redirecturl",
			"authurl", "tokenurl", "style")
	*/
	x := cmdtab.New("get")

	x.Usage = `<field> <name>`
	x.Summary = `print a field from the app configuration data`

	x.Description = `
		The *get* subcommand is usually invoked implicitly when a subcommand
		matches the name of a field in the application record. It can,
		however, also be used explicitly.`

	x.Method = func(args []string) error {
		if len(args) != 2 {
			return x.UsageError()
		}
		_, app, err := auth.Lookup(args[1])
		if err != nil {
			return err
		}
		switch args[0] {
		case "conf":
			fmt.Println(os.Getenv("AUTHCONF"))
		case "access":
			fmt.Println(app.AccessToken)
		case "refresh":
			fmt.Println(app.RefreshToken)
		case "type":
			fmt.Println(app.TokenType)
		case "expiry":
			fmt.Println(app.Expiry)
		case "state":
			fmt.Println(app.AuthState)
		case "code":
			fmt.Println(app.AuthCode)
		case "id":
			fmt.Println(app.ClientID)
		case "secret":
			fmt.Println(app.ClientSecret)
		case "scopes":
			for _, s := range app.Scopes {
				fmt.Println(s)
			}
		case "redirecturl":
			fmt.Println(app.RedirectURL)
		case "authurl":
			fmt.Println(app.Config.Endpoint.AuthURL)
		case "tokenurl":
			fmt.Println(app.Config.Endpoint.TokenURL)
		case "style":
			fmt.Println(app.Config.Endpoint.AuthStyle)
		}
		return nil
	}
}
