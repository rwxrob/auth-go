package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
)

// Data is a oath2-centric data structure designed to potentially
// hold configuration data for other authentication methods supported
// by this package. The oauth2.Config is embedded. This
// allows referencing different client applications by their
// user-friendly names.
type Data struct {
	oauth2.Config
	Name string
}

// Save writes the Data as compressed JSON (one line) into the file
// specified by path with the given permissions.
func (d *Data) Save(path string, perm os.FileMode) error {
	return ioutil.WriteFile(path, d.JSON(), perm)
}

// Parse is simply a wrapper for json.Unmarshal().
func (d *Data) Parse(buf []byte) error { return json.Unmarshal(buf, d) }

// Load loads the JSON data at path and unmarshals it with Parse().
func (d *Data) Load(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return d.Parse(buf)
}

// String implements the Stringer interface as long form JSON.
func (d Data) String() string {
	byt, _ := json.MarshalIndent(d, "", "  ")
	return string(byt)
}

// JSON returns a bytes buffer of compressed JSON suitable for saving
// and streaming. Otherwise it's the same as String().
func (d Data) JSON() []byte {
	byt, _ := json.Marshal(d)
	return byt
}

// Print prints as string to stdout.
func (d Data) Print() { fmt.Println(d) }
