package main

import (
	"fmt"

	"gitlab.com/rwxrob/auth"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("export")
	x.Usage = `<file> [<name>]`
	x.Summary = `writes one or all to JSON config file`

	x.Description = `
		If a name is provided writes only that specific auth application
		configuration data in JSON format to to the specified file path,
		which will be overritten if exists. If no name is provided writes
		the entire cache. Note that even if only one record is exported it
		will still be contained within a full config object in the resulting
		file. (For just one record without the full object use the *json*
		subcommand instead.)`

	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		config, err := auth.OpenConfig()
		if err != nil {
			return err
		}
		if len(args) > 1 {
			exp := auth.Config{}
			for _, name := range args[1:] {
				v, has := config[name]
				if !has {
					return fmt.Errorf("'%v' not found in cache", name)
				}
				if exp.Has(name) {
					return fmt.Errorf("'%v' passed twice", name)
				}
				exp[name] = v
			}
			return exp.Save(args[0])
		}
		return config.Save(args[0])
	}
}
