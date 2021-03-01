package main

import (
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New(
		"auth", "token", "grant", "ls", "get", "rm", "add", "import",
		"edit", "json", "refresh", "conf", "valid")
	x.Default = "token"
	x.Summary = `use and manage cached oauth2 and other authorizations`
	x.Version = "1.0.0"
	x.Author = "Rob Muhlestein <rwx@robs.io> (rwxrob.live)"
	x.Git = "github.com/rwxrob/auth/cmd/auth"
	x.Copyright = "(c) Rob Muhlestein"
	x.License = "Apache-2.0"

	x.Description = `
		The *auth* utility command is designed to make command line
		integration of Oauth2 and other standard authorizations easier and
		safer to manage. While the standards have been established for these
		important protocols the means of caching and managing the keys and
		tokens remains largely ad hoc. To address this problem a centralized
		store is used for authorization cache and configuration data from
		a single, secured JSON file not unlike other command line security
		tools such as ssh and gpg.

		The *token* subcommand is assumed if no other subcommands match.
		Therefore, do not use any app name identifiers that conflict with
		the listed subcommands.

		Interactivity

		The *auth* command is primary meant to be used from the command line
		by users directly or embedded into scripts that generally have user
		interaction. Therefore, attention should be given to which
		subcommands are used when creating automations that would be
		negatively impacted by blocking for user interaction. `

}
