package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
)

// Data is a oath2-centric data structure designed to potentially hold
// configuration data for other authentication methods supported by this
// package. The oauth2.Config is embedded as is an oauth2.Token. This
// allows referencing different client applications by their
// user-friendly names.
type Data struct {
	oauth2.Config
	oauth2.Token
	Name string
}

// Save writes the Data as compressed JSON (one line) into the file
// specified by path with the given permissions.
func (d *Data) Save(path string, perm os.FileMode) error {
	return ioutil.WriteFile(path, d.JSON(), perm)
}

// Save64 writes the Data as compressed JSON (one line) encoded into
// base 64 into the file specified by path with the given permissions.
func (d *Data) Save64(path string, perm os.FileMode) error {
	return ioutil.WriteFile(path, []byte(d.Base64()), perm)
}

// Parse is simply a wrapper for json.Unmarshal().
func (d *Data) Parse(buf []byte) error { return json.Unmarshal(buf, d) }

// Parse64 is as Parse but supports decoding base 64 JSON.
func (d *Data) Parse64(buf []byte) error {
	jsn, err := base64.StdEncoding.DecodeString(string(buf))
	if err != nil {
		return err
	}
	return json.Unmarshal(jsn, d)
}

// Load loads the JSON data at path and unmarshals it with Parse().
func (d *Data) Load(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return d.Parse(buf)
}

// Load64 loads the JSON data encoded in base64 at path and unmarshals it.
func (d *Data) Load64(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return d.Parse64(buf)
}

// TODO add SaveSecure()
// TODO add LoadSecure()
// TODO add Data.Encrypted()

// String implements the Stringer interface as long form JSON.
func (d Data) String() string {
	byt, _ := json.MarshalIndent(d, "", "  ")
	return string(byt)
}

// Print prints as string to stdout.
func (d Data) Print() { fmt.Println(d) }

// JSON returns a bytes buffer of compressed JSON suitable for saving
// and streaming. Otherwise it's the same as String().
func (d Data) JSON() []byte {
	byt, _ := json.Marshal(d)
	return byt
}

// Base64 returns a string containing base64 encoded JSON.
func (d Data) Base64() string {
	return base64.StdEncoding.EncodeToString(d.JSON())
}

// Mock returns an exact duplicate of itself but with all URLs,
// including the Endpoints, replaced with copies that have had their
// initial scheme, domain, and port replaced with the prefix passed to
// it. This allows any existing configuration data to be tested when
// combined with a local mock endpoint from StartMockEndpoint().
func (d Data) Mock(pre string) *Data {
	// TODO make sure there aren't alread Go packages that do this
	return nil
}
