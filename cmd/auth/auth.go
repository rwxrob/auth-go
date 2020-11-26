package main

import (
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New(
		"auth", "token", "grant", "ls", "get", "rm", "add", "import",
		"edit", "json", "refresh", "conf")
	x.Summary = `use and manage cached oauth2 and other authorizations`
	x.Version = "1.0.0"
	x.Author = "Rob Muhlestein <rwx@robs.io> (rwxrob.live)"
	x.Git = "gitlab.com/rwxrob/auth/cmd/auth"
	x.Copyright = "(c) Rob Muhlestein"
	x.License = "Apache-2.0"

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

}
