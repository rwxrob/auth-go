package main

import (
	"github.com/rwxrob/auth"
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("rm")
	x.Usage = `<name>`
	x.Summary = `remove an application from the local cache`

	x.Description = `
		The *rm* subcommand removes the specified entry from the app cache
		if it exists returning an error if it does not.`

	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}
		name := args[0]
		config, _, err := auth.Lookup(name)
		if err != nil {
			return err
		}
		delete(config, name)
		config.Store()
		return nil
	}
}
