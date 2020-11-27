package auth

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"gitlab.com/rwxrob/prompt"
	"golang.org/x/oauth2"
)

// Has returns 1 if the given name exists in the cache, 0 if not, and -1
// if cannot determine.
func Has(name string) int8 {
	config, err := OpenConfig()
	if err != nil {
		return -1
	}
	if config.Has(name) {
		return 1
	}
	return 0
}

// Valid returns one of four values: 1 if an access token exists for the
// given name and has not expired, 0 if a token exists but has expired
// or has no refresh token, -1 if there is no access token, and -2 if an
// error prevents determining if it exists.
func Valid(name string) int8 {
	switch Has(name) {
	case 0:
		return -1
	case -1:
		return -2
	}
	_, app, err := Lookup(name)
	if err != nil {
		return -2
	}
	switch app.Valid() {
	case true:
		return 1
	case false:
		return 0
	}
	return -2
}

// Lookup returns a Config loaded from the configuration file cache and
// a reference to the specified app if found. An error is also returned
// to explain if either of them are nil for any reason.
func Lookup(name string) (Config, *App, error) {
	config, err := OpenConfig()
	if err != nil {
		return nil, nil, err
	}
	app, has := config[name]
	if !has {
		return nil, nil, fmt.Errorf("'%v' not found", name)
	}
	return config, app, nil
}

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

// Authorize runs through the minimum required Oauth2 authorization flow
// avoiding interactive user input when possible by starting up a local
// HTTP server (or using the one that has already been started) to
// capture the incoming redirected data.
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
		code := prompt.Secret("Enter authorization code (echo off):")
		a.SetAuthCode(code)
		return nil
	}
	return nil
}
