package oauth

import (
	"fmt"
	"io/ioutil"
)

type AppData struct {
	Name         string `json:"app_name,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Code         string `json:"code,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	Expires      int    `json:"expires,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	Type         string `json:"token_type,omitempty"`
	LoginURI     string `json:"login_uri,omitempty"`
	TokenURI     string `json:"token_uri,omitempty"`
}

// String fulfills the Stringer interface as JSON
func (app AppData) String() string { return toJSON(app) }

// Print prints in JSON long form.
func (app *AppData) Print() {
	if app == nil {
		fmt.Println("null")
		return
	}
	fmt.Println(app)
}

// Save writes the JSON long form to the specified path.
func (app *AppData) Save(path string) error {
	return ioutil.WriteFile(path, []byte(toJSON(app)), 0600)
}
