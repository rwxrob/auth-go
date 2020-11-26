package main

import (
	"strconv"
	"strings"

	"gitlab.com/rwxrob/auth"
	"gitlab.com/rwxrob/cmdtab"
	"gitlab.com/rwxrob/prompt"
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
		app.Name, err = prompt.Until("Name (required): ")
		if err != nil {
			return err
		}
		app.ClientID, err = prompt.Until("ClientID (required): ")
		if err != nil {
			return err
		}
		app.ClientSecret, err = prompt.UntilSecret("ClientSecret (required): ")
		if err != nil {
			return err
		}
		app.RedirectURL, err = prompt.UntilStrict("RedirectURL (required): ", "^https?://")
		if err != nil {
			return err
		}
		app.Endpoint.AuthURL, err = prompt.UntilStrict("AuthURL (required): ", "^https?://")
		if err != nil {
			return err
		}
		app.Endpoint.TokenURL, err = prompt.UntilStrict("TokenURL (required): ", "^https?://")
		if err != nil {
			return err
		}
		style, err := prompt.UntilStrict("AuthStyle: ", "^[0-3]$")
		if err != nil {
			return err
		}
		stylei, err := strconv.Atoi(style)
		if err != nil {
			return err
		}
		app.Endpoint.AuthStyle = oauth2.AuthStyle(stylei)
		app.AccessToken, err = prompt.Plain("AccessToken: ")
		if err != nil {
			return err
		}
		app.RefreshToken, err = prompt.Plain("RefreshToken: ")
		if err != nil {
			return err
		}
		app.TokenType, err = prompt.Plain("TokenType: ")
		if err != nil {
			return err
		}
		app.AuthState, err = prompt.Plain("AuthState: ")
		if err != nil {
			return err
		}
		app.AuthCode, err = prompt.Plain("AuthCode: ")
		if err != nil {
			return err
		}
		scopes, err := prompt.Plain("Scopes (separated by spaces): ")
		if err != nil {
			return err
		}
		if scopes != "" {
			app.Scopes = []string{}
			for _, scope := range strings.Split(scopes, " ") {
				app.Scopes = append(app.Scopes, scope)
			}
		}
		conf[app.Name] = app
		conf.Cache()
		return nil
	}
}
