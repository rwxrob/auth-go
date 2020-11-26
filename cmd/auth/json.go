package main

import (
	"gitlab.com/rwxrob/auth"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("json")
	x.Usage = `<name>`
	x.Summary = `print JSON format of app to standard output`

	x.Description = `
			The *json* subcommand dumps the application cache for the
			specified app entry to standard output in long-form JSON format.`

	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}
		_, app, err := auth.Lookup(args[0])
		if err != nil {
			return err
		}
		app.Print()
		return nil
	}
}
