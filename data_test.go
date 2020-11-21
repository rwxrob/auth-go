package auth_test

import (
	"fmt"
	"testing"

	"gitlab.com/rwxrob/auth"
	"golang.org/x/oauth2"
)

func TestDataBase64(t *testing.T) {
	d := new(auth.Data)
	d.Name = "sample"
	d.ClientID = "client_id"
	d.ClientSecret = "client_secret"
	d.RedirectURL = "https://localhost:8080/callback"
	d.Scopes = []string{"some.scope", "another.scope"}
	d.Endpoint.AuthURL = "https://localhost:8081/oauth/authorize"
	d.Endpoint.TokenURL = "https://localhost:8081/oauth/token"
	d.Endpoint.AuthStyle = oauth2.AuthStyleInHeader
	t.Log(d.Base64())
}

func ExampleData_Load64() {
	d := new(auth.Data)
	d.Load64("testdata/sample")
	fmt.Println(
		d.Name,
		d.ClientID,
		d.ClientSecret,
		// etc.
	)

	// Output:
	// sample client_id client_secret

}
