package oauth

import (
	"fmt"
	"io/ioutil"
	"time"

	"gitlab.com/rwxrob/uniq"
)

type AppData struct {
	Name         string `json:"name,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Code         string `json:"code,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	Expires      int64  `json:"expires,omitempty"` // should be standard
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	Type         string `json:"token_type,omitempty"`
	State        string `json:"state,omitempty"`
	LoginURI     string `json:"login_uri,omitempty"`
	TokenURI     string `json:"token_uri,omitempty"`
}

// SetExpires sets the Expires value to the current time plus the
// ExpiresIn value.
func (d *AppData) SetExpires() {
	d.Expires = time.Now().Unix() + d.ExpiresIn
}

// SetState sets the State to a random base32 string (see
// gitlab.com/rwxrob/uniq).
func (d *AppData) SetState() { d.State = uniq.Base32() }

// TimeLeft returns the number of seconds before Expires
func (d *AppData) TimeLeft() int64 { return d.Expires - time.Now().Unix() }

// String fulfills the Stringer interface as JSON
func (d AppData) String() string { return toJSON(d) }

// Print prints in JSON long form.
func (d *AppData) Print() {
	if d == nil {
		fmt.Println("null")
		return
	}
	fmt.Println(d)
}

// Save writes the JSON long form to the specified path.
func (d *AppData) Save(path string) error {
	return ioutil.WriteFile(path, []byte(toJSON(d)), 0600)
}
