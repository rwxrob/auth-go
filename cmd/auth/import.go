package main

import (
	"fmt"

	"github.com/rwxrob/auth"
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("import")
	x.Usage = `<file> [<name>]...`
	x.Summary = `import an app config JSON file into cache`

	x.Description = `
		The *import* subcommand loads the specified <file> (in JSON Config
		cache format) into the local auth configuration cache overwriting
		anything with the same name. If one or more of <name> is provided in
		addition to <file> only those names will be imported.`

	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		imp := auth.Config{}
		err := imp.Load(args[0])
		if err != nil {
			return err
		}
		conf, err := auth.OpenConfig()
		if err != nil {
			return err
		}
		if len(args) > 1 {
			for _, name := range args[1:] {
				app, has := imp[name]
				if !has {
					return fmt.Errorf("'%v' not found in import", name)
				}
				conf[name] = app
			}
			return nil
		}
		for name, app := range imp {
			conf[name] = app
		}
		return conf.Store()
	}
}
