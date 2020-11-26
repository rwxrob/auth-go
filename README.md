# Go Oauth2 Helper Module

![Oauth2 Session](doc/session.gif)

[![GoDoc](https://godoc.org/gitlab.com/rwxrob/auth?status.svg)](https://godoc.org/gitlab.com/rwxrob/auth)
[![License](https://img.shields.io/badge/license-Apache-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/rwxrob/auth)](https://goreportcard.com/report/gitlab.com/rwxrob/auth)
[![Coverage](https://gocover.io/_badge/gitlab.com/rwxrob/auth)](https://gocover.io/gitlab.com/rwxrob/auth)

Designed to help make command line Oauth2 easier to implement and use
from shell scripts and other command line utilities.

## Example Usage

### Main

```
auth token <name>
auth grant <name>
auth refresh <name>
auth help
```

### Create, Replace, Update, Delete (CRUD)

```
auth add
auth rm <name>
auth edit
auth export <file> [<name> ...]
auth import <file> [<name> ...]
```

### View / Output

```
auth ls
auth conf 
auth get json <name>
auth get access <name>
auth get refresh <name>
auth get type <name>
auth get expiry <name>
auth get state <name>
auth get code <name>
auth get id <name>
auth get secret <name>
auth get scopes <name>
auth get redirecturl <name>
auth get authurl <name>
auth get tokenurl <name>
auth get style <name>
```

### Embed

```
curl -H "Authorization: Bearer $(auth token <name>)" https://api.example.com/some
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

* As much test coverage as we can achieve without a mock token server
* Add some detection of changes to the file since opened so don't
  overwrite with currently running process with the cache open (like vi)
