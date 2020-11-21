package auth

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"

	"golang.org/x/oauth2"
	"golang.org/x/term"
)

// Changing AuthDirEnv affects every subsequent call to most top-level
// functions in this package. Changing it is not safe for concurrency.
var DirEnvVar = `AUTHDIR`

func Dir() string {
	return os.Getenv(DirEnvVar)
}

// Authorize
func Authorize(d *Data) error {
	url := d.AuthCodeURL("state", oauth2.AccessTypeOffline)
	err := OpenBrowser(url)
	if err != nil {
		fmt.Printf("Visit the URL for the auth dialog: \n%s\n\n", url)
	}
	code, err := Prompt("Enter authorization code:")
	if err != nil {
		return fmt.Errorf("openlocalbrowser: unable to read authorization code %v", err)
	}
	ctx := context.Background()
	tok, err := d.Exchange(ctx, code)
	if err != nil {
		return err
	}

	d.Token = *tok
	Cache(d)

	return nil
}

// OpenBrowser opens the default (usually graphical) web browser on the
// current system. Currently only linux, windows, and darwin are
// supported.
func OpenBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows", "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("openbrowser: unsupported platform: %s",
			runtime.GOOS,
		)
	}
}

// Prompt simply prompts the user to enter text interactively.
func Prompt(s string) (string, error) {
	fmt.Printf("%v ", s)
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return "", err
	}
	return input, nil
}

// PromptSecret prompts the user to enter text interactively but does
// not echo what they are typing to the screen.
func PromptSecret(s string) (string, error) {
	fmt.Printf("%v ", s)
	input, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	return string(input), err
}

// Load returns a Data structure having loaded it from the AUTHDIR
// returning nil if either AUTHDIR is not set or the named data is not
// found. Data is stored unencrypted within files with names matching the
// Data.Name. These files should be considered interally encapsulated
// and therefore never directly accessed since the saved file format may
// change over time just as one would generally not directly access such
// files cached within a web browser.
func Load(name string) *Data {
	path := filepath.Join(Dir(), name)
	data := new(Data)
	err := data.Load64(path)
	if err != nil {
		return nil
	}
	return data
}

// Cache saves the Data into the AUTHDIR as base64 encoded JSON. This
// format of the caching should never be relied on, however, as there is
// no guarantee this format might not change in the future. Indeed, it
// is likely optionally support some form of flat-file encryption
// eventually.
func Cache(d *Data) error {
	return d.Save64(Path(d.Name), 0600)
}

// Exists returns true if cached data exists for the given name.
func Exists(name string) bool {
	_, err := os.Stat(Path(name))
	if err != nil {
		return false
	}
	return true
}

// Path returns the path to the cached file within Dir(). Usually this
// will be an absolute path, but that is determined by the value of
// DirAuthVar. Also see Dir().
func Path(name string) string {
	if len(name) == 0 {
		return ""
	}
	return filepath.Join(Dir(), name)
}

// Import reads a JSON or YAML file and returns a pointer to
// a configuration Data structure. A number of field names may be used
// including those from the Oauth2 RFC. Import will intelligently infer
// the type and structure of the imported data path. Files can be in
// JSON, YAML, TOML, XML or base64 encoded versions of any of them. The
// first runes of each will be read to determine the files type
// (JSON='{', YAML='---', TOML='[data]', XML='<data>',
// Base64=~^[A-Za-z0-9+/]) Files may contain multiple configuration data
// structures (usually the result of an Export()) or just one.
func Import(file string) *Data {
	d := new(Data)
	// TODO
	return d
}

// ImportMany takes a path to a directory or single file and imports all
// the importable configurations it can from that location by
// intelligently sensing what is there in order to make short work of
// combining and organizing multiple configurations. When importing
// a directory all files will be opened and attempted and will
// recursively descend into all subdirectories adding a dot (.) for
// every depth level to avoid name collisions.
func ImportMany(from string) *Data {
	// TODO
	return nil
}

/* TODO later
// EncryptWithTouchAuth takes a byte array and encrypts it after
// requiring a touchable physical authentication device using the FIDO2
// U2F protocol. The first device found will automatically be assumed.
func EncryptWithTouchAuth(b []byte) ([]byte, error) {
	var enc []byte
	var err error

	// get all the devices and pick the first one

	// create the temporary keys

	// prompt for touch and encrypt

	//
	enc = []byte("blah")
	return enc, err
}
*/
