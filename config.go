package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config contains a collection of App structures in a way that is
// designed to be stored as a single file to facilitate encoding,
// transfer, and encryption.
type Config map[string]*App

// String implements the Stringer interface as long form JSON.
func (c Config) String() string {
	byt, _ := json.MarshalIndent(c, "", "  ")
	return string(byt)
}

// Print prints as string to stdout.
func (c Config) Print() { fmt.Println(c) }

// JSON returns a bytes buffer of compressed JSON suitable for
// streaming. Otherwise it's the same as String().
func (c Config) JSON() []byte {
	byt, _ := json.Marshal(c)
	return byt
}

// Save writes the Config to disk. By default this is written to
// os.UserConfigDirectory under the auth directory with the file name
// config.json. If the AUTHCONF environment variable is set will save to
// that file instead.
func (c Config) Save() error {
	return ioutil.WriteFile(ConfigFilePath(), []byte(c.String()), 0600)
}

// Parse is simply a wrapper for json.Unmarshal().
func (c *Config) Parse(buf []byte) error { return json.Unmarshal(buf, c) }

// Load loads the JSON data from the ConfigFilePath path.
func (c *Config) Load() error {
	buf, err := ioutil.ReadFile(ConfigFilePath())
	if err != nil {
		return err
	}
	return c.Parse(buf)
}
