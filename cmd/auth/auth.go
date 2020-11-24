package main

import (
	"fmt"
	"os"

	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New(
		"auth", "token", "grant", "ls", "rm", "add", "import", "edit",
		"scopes", "conf", "json", "access", "refresh", "type", "expiry",
		"state", "code", "id", "secret", "redirecturl", "authurl")

	x.Summary = `use and manage cached oauth2 and other authorizations`

	x.Description = `
		The *auth* utility command is designed to make command line
		integration of Oauth2 and other standard authorizations easier and
		safer to manage. While the standards have been established for these
		important protocols the means of caching and managing the keys and
		tokens remains largely ad hoc. The *auth* utility seeks to address
		this need by managing all authorization cache and configuration data
		from a single, secured JSON store not unlike other command line
		security tools such as ssh and gpg.

		Interactivity

		The *auth* command is primary meant to be used from the command line
		by users directly or embedded into scripts that generally have user
		interaction. Therefore, attention should be given to which
		subcommands are used when creating automations that would be
		negatively impacted by blocking for user interaction. `

	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		switch args[0] {
		case "access", "refresh", "expiry", "state", "code",
			"id", "secret", "scopes", "redirecturl", "authurl",
			"tokenurl", "style":
			return cmdtab.Call("get", args)
		case "conf":
			fmt.Println(os.Getenv("AUTHCONF"))
			return nil
		default:
			if cmdtab.Has(args[0]) {
				return cmdtab.Call(args[0], args[1:])
			}
			return cmdtab.Call("token", args[1:])
		}
	}
}
