package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Token struct {
	Access    string `json:"access_token"`
	Expires   int    `json:"expires,omitempty"`
	ExpiresIn int    `json:"expires_in,omitempty"`
	Refresh   string `json:"refresh_token,omitempty"`
	Scope     string `json:"scope,omitempty"`
	Type      string `json:"token_type,omitempty"`
}

// Print prints in JSON long form.
func (t *Token) Print() {
	if t == nil {
		fmt.Println("null")
		return
	}
	fmt.Println(t)
}

// String fulfills the Stringer interface as JSON
func (t Token) String() string { return toJSON(t) }

// HaveToken returns true if token data for an app has been cached.
func HaveToken(app string) bool {
	return exists(path(app, "token.json"))
}

// LoadToken returns pointer to the token data for an app if it exists
// or nil otherwise.
func LoadToken(app string) *Token {
	buf, err := ioutil.ReadFile(path(app, "token.json"))
	if err != nil {
		log.Print(err)
		return nil
	}
	tk := new(Token)
	err = json.Unmarshal(buf, tk)
	if err != nil {
		log.Print(err)
		return nil
	}
	return tk
}

// CacheToken takes the usual JSON returned when a new token is created.
func CacheToken(jsn []byte) error {
	// TODO
	return nil
}
