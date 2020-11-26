package main

import (
	"fmt"
	"os"

	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("conf")
	x.Summary = `print the path the cache file (AUTHCONF)`

	x.Description = `
		The *conf* subcommand simply prints the content of the AUTHCONF
		environment variable.`

	x.Method = func(args []string) error {
		fmt.Println(os.Getenv("AUTHCONF"))
		return nil
	}
}
