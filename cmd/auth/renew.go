package main

import (
	"fmt"

	"gitlab.com/rwxrob/auth"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("renew")
	x.Usage = `<appname>`
	x.Summary = `renew an application oauth2 authorization`
	x.Description = `
		The *renew* subcommand discards all cached Oauth2 token data for the
		specified application and initiates a full authorization flow using
		the grant the *code* grant type. This includes creating a temporary
		HTTP server to handle the redirected URI capturing the code without
		any interaction from the user other than authenticating through
		their web browser.
	`
	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		c, err := auth.LoadConfig()
		if err != nil {
			return err
		}
		if a, found := c[args[0]]; found {
			err := a.Authorize()
			if err != nil {
				return err
			}
			c.Save()
		}
		return fmt.Errorf("no auth entry for '%v'", args[0])
	}
}
