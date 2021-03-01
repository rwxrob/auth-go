package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/rwxrob/auth"
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("edit")
	x.Summary = `open the configuration file directly for editing`

	x.Description = `
		Edit the configuration file (auth conf) directly with the default
		editor. If the EDITOR environment variable is set that will be
		opened. Otherwise, the system default file opener will be used (ex:
		xdg-open, open).`

	x.Method = func(args []string) error {
		if len(args) > 0 {
			return x.UsageError()
		}
		conf := os.Getenv("AUTHCONF")
		if conf == "" {
			return errors.New("AUTHCONF undefined")
		}
		cmd := os.Getenv("EDITOR")
		if cmd == "" {
			return auth.OpenResource(conf)
		}
		ed := exec.Command(cmd, conf)
		ed.Stdin = os.Stdin
		ed.Stdout = os.Stdout
		ed.Run()
		return nil
	}
}
