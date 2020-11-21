package auth_test

import (
	"os"
	"testing"

	"gitlab.com/rwxrob/auth"
	"golang.org/x/oauth2"
)

func ExampleData() {

	// initializing
	d := new(auth.Data)
	d.Name = "sample"
	d.ClientID = "client_id"
	d.ClientSecret = "client_secret"
	d.RedirectURL = "https://localhost:8080/callback"
	d.Scopes = []string{"some.scope", "another.scope"}
	d.Endpoint.AuthURL = "https://localhost:8081/oauth/authorize"
	d.Endpoint.TokenURL = "https://localhost:8081/oauth/token"
	d.Endpoint.AuthStyle = oauth2.AuthStyleInHeader

	// AccessToken and RefreshToken are only set from a successful auth

	// saving
	d.Save("something.json", 0600)

	// loading
	another := new(auth.Data)
	another.Load("testdata/something.json")
}

func ExampleLoad() {
	os.Setenv(auth.DirEnvVar, "testdata")
	d := auth.Load("sample")
	if d != nil {
		d.Print()
	}
	// Output:
	// {
	//   "ClientID": "client_id",
	//   "ClientSecret": "client_secret",
	//   "Endpoint": {
	//     "AuthURL": "https://localhost:8081/oauth/authorize",
	//     "TokenURL": "https://localhost:8081/oauth/token",
	//     "AuthStyle": 2
	//   },
	//   "RedirectURL": "https://localhost:8080/callback",
	//   "Scopes": [
	//     "some.scope",
	//     "another.scope"
	//   ],
	//   "access_token": "",
	//   "expiry": "0001-01-01T00:00:00Z",
	//   "Name": "sample"
	// }
}

func ExamplePrompt() {
	val, err := auth.Prompt("Enter something:")
	_ = val
	_ = err
}

func ExamplePromptSecret() {
	sec, err := auth.PromptSecret("Enter something secret:")
	_ = sec
	_ = err
}

func TestEnvRestored(t *testing.T) {
	t.Log(os.Getenv(auth.DirEnvVar))
}
