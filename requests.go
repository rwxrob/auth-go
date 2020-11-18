package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RequestTokenFromCode assumes the AppData contains a short-lived Code
// (grant_type=authorization_code) that can be promoted into a full
// access_token (and accompanying refresh_token). The JSON data returned
// form the request is unmarshaled directly into the AppData struct
// passed as a pointer in a way that is safe for concurrency.
func RequestTokenFromCode(d *AppData) error {
	var err error

	// validate required stuff
	switch {
	case d.TokenURI == "":
		return fmt.Errorf("TokenURI required")
	case d.ClientId == "":
		return fmt.Errorf("ClientId required")
	case d.ClientSecret == "":
		return fmt.Errorf("ClientSecret required")
	case d.RedirectURI == "":
		return fmt.Errorf("RedirectURI required")
	case d.Code == "":
		return fmt.Errorf("Code required")
	}

	// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --user [your client id]:[your client secret] --data "grant_type=authorization_code&redirect_uri=[your redirect URI]&code=[code]" https://api.restream.io/oauth/token

	v := url.Values{}
	v.Set("grant_type", "authorization_code")
	v.Set("redirect_uri", d.RedirectURI)
	v.Set("code", d.Code)

	// needs 'long' form for SetBasicAuth
	req, err := http.NewRequest(
		"POST", d.TokenURI,
		strings.NewReader(v.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(d.ClientId, d.ClientSecret)

	// send it
	res, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}

	// extract the tokens, expiration, and such
	defer res.Body.Close()
	byt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	d.Lock()
	defer d.Unlock()
	err = json.Unmarshal(byt, d)
	if err != nil {
		return err
	}

	return nil
}
