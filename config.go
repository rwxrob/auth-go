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

// Has returns true of the name exists.
func (c Config) Has(name string) bool {
	_, has := c[name]
	return has
}

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

// Save writes the config JSON data to a file at specified path.
func (c Config) Save(path string) error {
	return ioutil.WriteFile(path, []byte(c.String()), 0600)
}

// Store saves to the file specified by ConfigFilePath(). See Save().
func (c Config) Store() error { return c.Save(ConfigFilePath()) }

// Parse is simply a wrapper for json.Unmarshal().
func (c *Config) Parse(buf []byte) error { return json.Unmarshal(buf, c) }

// Open loads the JSON data from the ConfigFilePath path.
func (c *Config) Open() error { return c.Load(ConfigFilePath()) }

// Load reads the JSON data from the specified path.
func (c *Config) Load(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return c.Parse(buf)
}
