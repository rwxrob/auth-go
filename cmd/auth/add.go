package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rwxrob/auth-go"
	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/prompt-go"
	"golang.org/x/oauth2"
)

func init() {
	x := cmdtab.New("add")
	x.Usage = ``
	x.Summary = `interactively add a new application`

	x.Description = `
		The *add* subcommand starts an interactive prompt to create a new
		entry in the application cache prompting for each field. All fields
		can be skipped by entering blank values and filled in later with the
		*edit* subcommand.`

	x.Method = func(args []string) (err error) {
		conf, err := auth.OpenConfig()
		if err != nil {
			return err
		}
		app := new(auth.App)
		app.Name = prompt.Until("Name (required): ")
		app.ClientID = prompt.Plain("ClientID: ")
		app.ClientSecret = prompt.Plain("ClientSecret: ")
		app.Endpoint.AuthURL = prompt.Plain("AuthURL: ")
		app.Endpoint.TokenURL = prompt.Plain("TokenURL: ")
		style := prompt.Plain("AuthStyle: ")
		stylei, _ := strconv.Atoi(style)
		app.Endpoint.AuthStyle = oauth2.AuthStyle(stylei)
		app.AccessToken = prompt.Plain("AccessToken: ")
		app.RefreshToken = prompt.Plain("RefreshToken: ")
		app.TokenType = prompt.Plain("TokenType: ")
		app.AuthState = prompt.Plain("AuthState: ")
		app.AuthCode = prompt.Plain("AuthCode: ")
		scopes := prompt.Plain("Scopes (separated by spaces): ")
		if scopes != "" {
			app.Scopes = []string{}
			for _, scope := range strings.Split(scopes, " ") {
				app.Scopes = append(app.Scopes, scope)
			}
		}
		app.RedirectURL = "http://localhost:8080/redirected"
		fmt.Println("Remember to add expected redirect URL:\n  " + app.RedirectURL)
		conf[app.Name] = app
		conf.Store()
		return nil
	}
}
