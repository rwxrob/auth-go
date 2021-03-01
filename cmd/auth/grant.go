package main

import (
	"github.com/rwxrob/auth"
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("grant")
	x.Usage = `<name>`
	x.Summary = `grant an single application oauth2 authorization`

	x.Description = `
		The *grant* subcommand discards all cached Oauth2 token data for the
		specified application and initiates a full authorization flow using
		the *code* grant type. This includes creating a temporary HTTP
		server to handle the redirected URI capturing the code without any
		interaction from the user other than authorizing the application as
		they normally would through their web browser. Note that this
		ignores the expiry of a token if it has already been granted.
		Consider the (default) *token* subcommand instead for such
		requirements.`

	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		return auth.Grant(args[0])
	}
}
