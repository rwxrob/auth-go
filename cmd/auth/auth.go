package main

import (
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("auth", "token", "renew")
	//x.Default = "token"
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
	`
}
