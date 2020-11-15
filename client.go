package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Client struct {
	Id     string `json:"client_id,omitempty"`
	Secret string `json:"client_secret,omitempty"`
	Code   string `json:"code,omitempty"`
	Token  *Token `json:"token,omitempty"`
}

// String fulfills the Stringer interface as JSON
func (c Client) String() string { return toJSON(c) }

// Print prints in JSON long form.
func (c *Client) Print() {
	if c == nil {
		fmt.Println("null")
		return
	}
	fmt.Println(c)
}

// HaveClient returns true if client data for an app has been cached.
func HaveClient(app string) bool {
	return exists(path(app, "client.json"))
}

// LoadClient returns pointer to the client data for an app if it exists
// or nil otherwise.
func LoadClient(app string) *Client {
	buf, err := ioutil.ReadFile(path(app, "client.json"))
	if err != nil {
		log.Print(err)
		return nil
	}
	c := new(Client)
	err = json.Unmarshal(buf, c)
	if err != nil {
		log.Print(err)
		return nil
	}
	if HaveToken(app) {
		c.Token = LoadToken(app)
	}
	return c
}
