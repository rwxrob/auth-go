package oauth

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Update is the main entry to the oauth authorization flow for an
// existing application (one that has already been added with Add().
func Update(app string) error {
	// TODO
	return nil
}

// Cache updates the cache file for the given app and returns itself.
// Any changes to the underlying cache files will be overwritten.
func Cache(app *AppData) error {
	return app.Save(path(app.Name + ".json"))
}

// Have returns true if data for an app has been cached.
func Have(app string) bool {
	return exists(path(app + ".json"))
}

// Load returns pointer to the application oauth data if it exists or
// nil otherwise.
func Load(app string) *AppData {
	buf, err := ioutil.ReadFile(path(app + ".json"))
	if err != nil {
		log.Print(err)
		return nil
	}
	appdata := new(AppData)
	err = json.Unmarshal(buf, appdata)
	if err != nil {
		log.Print(err)
		return nil
	}
	return appdata
}
