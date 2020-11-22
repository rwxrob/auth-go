package main

import (
	"fmt"
	"time"

	"gitlab.com/rwxrob/auth"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("renew")
	x.Usage = `<appname>`
	x.Summary = `renew an single application oauth2 authorization`
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
			// just bail if token hasn't expired
			if a.RefreshToken != "" && !a.Expiry.Before(time.Now()) {
				return nil
			}
			exp := a.Expiry
			err := auth.Authorize(a)
			if err != nil {
				return err
			}
			fmt.Println("Wait for authorization to complete.")
			fmt.Println("(Cancel with Ctrl-C if necessary.)")
			// FIXME not quite there yet
			for {
				if a.Expiry != exp {
					break
				}
				time.Sleep(300 * time.Millisecond)
			}
			c.Save()
		}
		return fmt.Errorf("no auth entry for '%v'", args[0])
	}
}
