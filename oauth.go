package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

// Update is the main entry to the oauth authorization flow for an
// existing cached oauth application.
func Update(app string) error {

	// load the AppData and check for client_id and client_secret
	d := Load(app)
	switch {
	case d == nil:
		return fmt.Errorf("update: failed to load app data")
	case d.ClientId == "":
		return fmt.Errorf("update: client_id not set")
	case d.ClientSecret == "":
		return fmt.Errorf("update: client_secret not set")
	}

	// if no expires then set
	if d.Expires == 0 {
		d.Expires = time.Now().Unix()
		Cache(d)
	}

	// if access_token has more than 10 minutes of life left just return
	if d.TimeLeft() < 600 {
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
	fmt.Println("would refresh")
	return nil
}

func authorize(d *AppData) error {
	// TODO start local redirect server that will upgrade code into token
	// TODO open local browser to login uri
	fmt.Println("would initialize")
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
	buf, err := ioutil.ReadFile(path(app + ".json"))
	if err != nil {
		log.Print(err)
		return nil
	}
	return Parse(buf)
}

// Parse takes a byte buffer of JSON data, parses it, and returns
// pointer to the application oauth data if it exists or nil otherwise.
func Parse(buf []byte) *AppData {
	appdata := new(AppData)
	err := json.Unmarshal(buf, appdata)
	if err != nil {
		log.Print(err)
		return nil
	}
	return appdata
}

// OpenLocalBrowser opens the users local graphic browser (with
// xdg-open) TODO: detect other platforms
func OpenLocalBrowser(url string) {
	exec.Command("xdg-open", url).Start()
}
