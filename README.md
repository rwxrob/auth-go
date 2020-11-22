# Go Oauth2 Helper Module

[![GoDoc](https://godoc.org/gitlab.com/rwxrob/auth-go?status.svg)](https://godoc.org/gitlab.com/rwxrob/auth-go)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/rwxrob/auth-go)](https://goreportcard.com/report/gitlab.com/rwxrob/auth-go)
[![Coverage](https://gocover.io/_badge/gitlab.com/rwxrob/auth-go)](https://gocover.io/gitlab.com/rwxrob/auth-go)

Designed to help make command line Oauth2 easier to implement and use
from shell scripts and other command line utilities.

## Command Usage

```
auth token
```

### General Management

When run without any arguments `auth` enters interactive REPL user mode
prompting for actions and input.

All actions can be executed without interaction by entering the
subcommands and arguments:

```sh
auth create gitlab-rwxrob
auth import gitlab-rwxrob.yml
```

### Shell Script Integration

The `auth` command can be used in place of sensitive token and other
credential information within shell scripts. By default, the user will
be prompted when further authorization flow steps are needed, including
the opening of a local graphical web browser automatically. The level of
interaction can be isolated for scripts that must run without blocking
on waits for interactivity.

```sh
curl ... $(auth token gitlab)
```

By default, if a token of the given type is not cached or has expired
the user is prompted to authorize a new one.

## TODO

* Integrate the main 3-leg Oauth2 flow
* Mock endpoint for testing
