package auth

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

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

// Use is a the highest level entry point to this package. It returns
// the same values as Lookup() but also does whatever work is necessary
// to ensure that the named application has an updated access token.
// This includes potentially triggering the interactive flow with the
// user requiring authentication to the application through their web
// browser. For this reason Use() should not be called in situations
// where blocking on such interaction is not wanted. In such cases use
// Lookup() instead. Note that the named application should already
// exist and be present in the file located at ConfigFilePath().
func Use(name string) (Config, *App, error) {
	config, app, err := Lookup(name)
	if err != nil {
		return nil, nil, err
	}
	err = app.Refresh()
	if err == nil {
		config.Store()
		return config, app, nil
	}
	err = Grant(app) // includes config.Store
	return config, app, nil
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

// Grant runs through the Oauth2 authorization code grant flow avoiding
// interactive user input when possible by starting up a local HTTP
// server (or using the one that has already been started) to capture
// the incoming redirected data.
func Grant(this interface{}) error {
	var a *App
	config, err := OpenConfig()

	// prompt arg into app ref
	switch v := this.(type) {
	case *App:
		a = v
	case string:
		a = config[v]
	}
	exp := a.Expiry

	// start server and redirect app to it to get auth code
	AddSession(a)
	StartLocalServer()

	// open the user's web browser or prompt for auth code
	fmt.Println("Attempting to open your web browser")
	url := a.AuthCodeURL(a.AuthState, oauth2.AccessTypeOffline)
	err = OpenResource(url)
	if err != nil {
		fmt.Printf("Visit the URL for the auth dialog: \n  %s\n\n", url)
		code := prompt.Secret("Enter authorization code (echo off):")
		a.SetAuthCode(code)
		return nil
	}

	// wait for redirect data update the app
	fmt.Println("Wait for authorization to complete.")
	fmt.Println("(Cancel with Ctrl-C if necessary.)")
	for {
		if a.Expiry != exp {
			break
		}
		time.Sleep(300 * time.Millisecond)
	}

	return config.Store()
}
