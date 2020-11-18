package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

const (
	DefaultLoginURI    = `http://localhost:8080/login`
	DefaultRedirectURI = `http://localhost:8080/oauth`
	DefaultTokenURI    = `http://localhost:8080/token`
)

// Update is the main entry to the oauth authorization flow for an
// existing cached oauth application. The argument passed can either be
// the unique name from the cached app data or a pointer to an AppData
// struct.
func Update(app interface{}) error {
	var d *AppData

	switch v := app.(type) {
	case string:
		d = Load(v)
	case *AppData:
		d = v
	}

	// load the AppData and check for client_id and client_secret
	switch {
	case d == nil:
		return fmt.Errorf("update: failed to load app data")
	case d.ClientId == "":
		return fmt.Errorf("update: client_id not set")
	case d.ClientSecret == "":
		return fmt.Errorf("update: client_secret not set")
	}

	//set values to their defaults so it becomes obvious how they can be
	//updated in the cache files
	if d.LoginURI == "" {
		d.LoginURI = DefaultLoginURI
	}
	if d.RedirectURI == "" {
		d.RedirectURI = DefaultRedirectURI
	}
	if d.TokenURI == "" {
		d.TokenURI = DefaultTokenURI
	}
	if d.Expires == 0 {
		d.Expires = time.Now().Unix()
		Cache(d)
	}
	// if access_token has more than 10 minutes of life left just return
	if d.TimeLeft() > 600 {
		return nil
	}

	// if refresh_token then refresh request, if fails continue
	if d.RefreshToken != "" {
		err := refresh(d)
		if err == nil {
			return nil
		}
	}

	// request a brand new authorization token through graphic browser
	return authorize(d)

}

func refresh(d *AppData) error {
	// TODO
	fmt.Println("would refresh")
	return nil
}

func authorize(d *AppData) error {

	// set the state and url for this authorization flow
	d.SetState()
	u := "%v?response_type=code&client_id=%v&redirect_uri=%v&state=%v"
	url := fmt.Sprintf(u, d.LoginURI, d.ClientId, d.RedirectURI, d.State)

	// open a local http server to handle the incoming redirect data
	com := make(chan interface{})
	// FIXME needs to start *another* server
	fmt.Println(d)
	if err := StartServer(d, d.RedirectURI, com); err != nil {
		return err
	}

	// fire off the local graphic user browser for them to login
	OpenLocalBrowser(url)

	// wait around for the right data
	<-com

	return nil
}

// Cache updates the cache file for the given app and returns itself.
// Any changes to the underlying cache files will be overwritten.
func Cache(d *AppData) error {
	return d.Save(path(d.Name + ".json"))
}

// Have returns true if data for an app has been cached.
func Have(name string) bool {
	return exists(path(name + ".json"))
}

// Load returns pointer to the application oauth data if it exists or
// nil otherwise.
func Load(app string) *AppData {
	buf, _ := ioutil.ReadFile(path(app + ".json"))
	return Parse(buf)
}

// Parse takes a byte buffer of JSON data, parses it, and returns
// pointer to the application oauth data if it exists or nil otherwise.
func Parse(buf []byte) *AppData {
	appdata := new(AppData)
	json.Unmarshal(buf, appdata)
	return appdata
}

// OpenLocalBrowser opens the users local graphic browser (with
// xdg-open) TODO: detect other platforms
func OpenLocalBrowser(url string) {
	exec.Command("xdg-open", url).Start()
}
