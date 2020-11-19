package auth_test

import (
	"gitlab.com/rwxrob/auth"
	"golang.org/x/oauth2"
)

func ExampleData_String() {

	d := new(auth.Data)
	d.Name = "sample"
	d.ClientID = "client_id"
	d.ClientSecret = "client_secret"
	d.RedirectURL = "https://localhost:8080/callback"
	d.Scopes = []string{"some.scope", "another.scope"}
	d.Endpoint.AuthURL = "https://localhost:8081/oauth/authorize"
	d.Endpoint.TokenURL = "https://localhost:8081/oauth/token"
	d.Endpoint.AuthStyle = oauth2.AuthStyleInHeader
	d.Print() // same as fmt.Println(d)

	// Unordered Output:
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
	//   "Name": "sample"
	// }

}

func ExampleData() {
	d := new(auth.Data)
	d.Name = "something"
	d.Save("something.json", 0600)
	another := new(auth.Data)
	another.Load("testdata/something.json")
}
