module gitlab.com/rwxrob/auth

go 1.15

replace gitlab.com/rwxrob/cmdtab => ../cmdtab

require (
	gitlab.com/rwxrob/cmdtab v0.0.0-00010101000000-000000000000
	gitlab.com/rwxrob/uniq v0.0.0-20200325203910-f771e6779384
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58
	golang.org/x/term v0.0.0-20201117132131-f5c789dd3221
	golang.org/x/tools v0.0.0-20201121010211-780cb80bd7fb // indirect
)
