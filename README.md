# Go Oauth2 Helper Module

[![GoDoc](https://godoc.org/gitlab.com/rwxrob/auth-go?status.svg)](https://godoc.org/gitlab.com/rwxrob/auth-go)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/rwxrob/auth-go)](https://goreportcard.com/report/gitlab.com/rwxrob/auth-go)
[![Coverage](https://gocover.io/_badge/gitlab.com/rwxrob/auth-go)](https://gocover.io/gitlab.com/rwxrob/auth-go)

Designed to help make command line Oauth2 easier to implement and use
from shell scripts and other command line utilities.

## Example Usage

### Create, Replace, Update, Delete (CRUD)

```
auth import <file> (JSON,YAML,TOML,XML)
auth export <file>
auth add <name>
auth rm <name>
auth token <name>
auth grant <name>
auth refresh <name>
auth edit
```

### View / Output

```
auth ls
auth json <name>
auth yaml <name>
auth toml <name>
auth xml <name>
```

### Fields

```
auth access <name> (default)
auth refresh <name>
auth type <name>
auth expiry <name>
auth state <name>
auth code <name>
auth id <name>
auth secret <name>
auth scopes <name>
auth redirecturl <name>
auth authurl <name>
auth tokenurl <name>
auth style <name>
auth conf 
```
### Embed

```
curl -H "Authorization: Bearer $(auth <name>)" https://api.example.com/some
```

The `auth` command can be used in place of sensitive token and other
credential information within shell scripts. By default, the user will
be prompted when further authorization flow steps are needed, including
the opening of a local graphical web browser automatically. The level of
interaction can be isolated for scripts that must run without blocking
on waits for interactivity.

By default, if a token of the given type is not cached or has expired
the user is prompted to authorize a new one.

## TODO

* Need to polish up the `auth` command and test `App.Client()`
* Mock endpoint for testing (someday)
