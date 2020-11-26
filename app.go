package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"sync"

	"gitlab.com/rwxrob/uniq"
	"golang.org/x/oauth2"
)

// App is an oath2-centric data structure designed to potentially hold
// configuration data for other authentication methods supported by this
// package. The oauth2.Config is embedded as is oauth2.Token. This
// allows referencing different client applications by their
// user-friendly names while still using App exactly as would be done
// with either struct by itself.
type App struct {
	Name      string
	AuthState string
	AuthCode  string
	sync.RWMutex
	oauth2.Config
	oauth2.Token
}

// String implements the Stringer interface as long form JSON.
func (a App) String() string {
	byt, _ := json.MarshalIndent(a, "", "  ")
	return string(byt)
}

// Print prints as string to stdout.
func (a App) Print() { fmt.Println(a) }

// JSON returns a bytes buffer of compressed JSON suitable for saving
// and streaming. Otherwise it's the same as String().
func (a App) JSON() []byte {
	byt, _ := json.Marshal(a)
	return byt
}

// Save writes the App to disk at the specified path.
func (a App) Save(path string) error {
	return ioutil.WriteFile(path, []byte(a.String()), 0600)
}

// Parse is simply a wrapper for json.Unmarshal().
func (a *App) Parse(buf []byte) error { return json.Unmarshal(buf, a) }

// Load loads the JSON data from the specified path.
func (a *App) Load(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return a.Parse(buf)
}

// SetAuthState updates the state to a new unique (base32) string.
func (a *App) SetAuthState() {
	defer a.Unlock()
	a.Lock()
	a.AuthState = uniq.Base32()
}

// SetAuthCode updates the AuthCode safely.
func (a *App) SetAuthCode(code string) {
	defer a.Unlock()
	a.Lock()
	a.AuthCode = code
}

// ParseRedirectURI calls ParseRequestURI on RedirectURI.
func (a App) ParseRedirectURL() (*url.URL, error) {
	return url.ParseRequestURI(a.RedirectURL)
}

// RedirectHost returns just the Host and Port portions of the
// RedirectURL suitable for passing to ListenAndServe() (addr) when
// starting the local server.
func (a App) RedirectHost() string {
	url, _ := a.ParseRedirectURL()
	return url.Host
}

// Refresh submits the refresh_token grant request to the app.TokenURL.
func (a *App) Refresh() error {
	ts := a.TokenSource(oauth2.NoContext, &a.Token)
	nw, err := ts.Token()
	if err != nil {
		return err
	}
	if nw.AccessToken != a.AccessToken {
		a.Lock()
		a.Token = *nw
		a.Unlock()
	}
	return nil
}
