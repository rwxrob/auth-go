# Go Oauth2 Helper Module

[![Oauth2 Session](doc/session.gif)]

[![GoDoc](https://godoc.org/gitlab.com/rwxrob/auth-go?status.svg)](https://godoc.org/gitlab.com/rwxrob/auth-go)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/rwxrob/auth-go)](https://goreportcard.com/report/gitlab.com/rwxrob/auth-go)
[![Coverage](https://gocover.io/_badge/gitlab.com/rwxrob/auth-go)](https://gocover.io/gitlab.com/rwxrob/auth-go)

Designed to help make command line Oauth2 easier to implement and use
from shell scripts and other command line utilities.

## Example Usage

### Create, Replace, Update, Delete (CRUD)

```
auth token <name> (default)

auth refresh <name>

auth add
auth export <file> [<name> ...]
auth import <file> [<name> ...]
auth rm <name>
auth edit
auth grant <name>
```

### View / Output

```
auth conf 
auth ls
auth json <name>
auth access <name>
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

* Add some detection of changes to the file since opened so don't
  overwrite with currently running process with the cache open (like vi)
* Need to polish up the `auth` command and test `App.Client()`
* Mock endpoint for testing (someday)
