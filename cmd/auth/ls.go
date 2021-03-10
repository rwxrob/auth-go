package main

import (
	"fmt"

	"github.com/rwxrob/auth-go"
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("ls")
	x.Summary = `list the cached application data by name`

	x.Description = `
		The *ls* subcommand lists the names of every cached app record one
		to a line.`

	x.Method = func(args []string) error {
		c, err := auth.OpenConfig()
		if err != nil {
			return err
		}
		for k, _ := range c {
			fmt.Println(k)
		}
		return nil
	}
}
