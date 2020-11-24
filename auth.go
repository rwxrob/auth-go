package auth

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"

	"golang.org/x/oauth2"
	"golang.org/x/term"
)

// OpenConfig loads the configuration file (see Config). Returns nil if
// unable to load.
func OpenConfig() (Config, error) {
	c := new(Config)
	err := c.Open()
	if err != nil {
		return nil, err
	}
	return *c, nil
}

// ConfigFilePath returns the path to the configuration file. First the
// AUTHCONF env var is checked and if not found the os.UserConfigDir
// will be checked for an auth/config.json file.
func ConfigFilePath() string {
	file := os.Getenv("AUTHCONF")
	if file == "" {
		dir, err := os.UserConfigDir()
		if err != nil {
			return ""
		}
		dir = filepath.Join(dir, "auth")
		err = os.Mkdir(dir, 0700)
		if err != nil {
			return ""
		}
		file = filepath.Join(dir, "config.json")
	}
	return file
}

// OpenResource opens the the specified resource (URL, file, etc.) using
// the opener of the current system. Currently only linux, windows, and
// darwin are supported.
func OpenResource(res string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", res).Start()
	case "windows", "darwin":
		return exec.Command("open", res).Start()
	default:
		return fmt.Errorf("openresource: unsupported platform: %s",
			runtime.GOOS,
		)
	}
}

// Prompt simply prompts the user to enter text interactively.
func Prompt(s string) (string, error) {
	fmt.Printf("%v ", s)
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return "", err
	}
	return input, nil
}

// PromptSecret prompts the user to enter text interactively but does
// not echo what they are typing to the screen.
func PromptSecret(s string) (string, error) {
	fmt.Printf("%v ", s)
	input, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	return string(input), err
}

// Authorize runs through the minimum required Oauth2 authorization flow
// avoiding interactive user input when possible by starting up a local
// HTTP server (or using the one that has already been started) to
// capture the incoming redirected data. Note that the oauth2 package
// does not provide any way of detecting expired refresh tokens (only
// access tokens). Currently the best alternative is to trap
// oath2.Client() requests that are unsuccessful and call Authorize in
// such cases in addition to depending on the "automatic" token
// refreshing done by the oauth2.Client.
func Authorize(a *App) error {

	// start server and send app to it for caching
	AddSession(a)
	StartLocalServer()

	// open the user's web browser or prompt for auth code
	fmt.Println("Attempting to open your web browser")
	url := a.AuthCodeURL(a.AuthState, oauth2.AccessTypeOffline)
	err := OpenResource(url)
	if err != nil {
		fmt.Printf("Visit the URL for the auth dialog: \n  %s\n\n", url)
		code, err := PromptSecret("Enter authorization code (echo off):")
		if err != nil {
			return err
		}
		a.SetAuthCode(code)
		return nil
	}
	return nil
}
