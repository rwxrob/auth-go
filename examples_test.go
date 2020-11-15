package oauth_test

import (
	"fmt"
	"os"

	"gitlab.com/rwxrob/oauth"
)

func ExampleLoadClient() {
	os.Setenv("OAUTHDIR", "testdata")
	client := oauth.LoadClient("some.io")
	client.Print()
	none := oauth.LoadClient("none")
	none.Print()
	broke := oauth.LoadClient("broke")
	broke.Print()
	// Output:
	// {
	//   "client_id": "some_client_id",
	//   "client_secret": "some_client_secret"
	// }
	// null
	// null
}

func ExampleLoadToken() {
	os.Setenv("OAUTHDIR", "testdata")
	token := oauth.LoadToken("some.io")
	token.Print()
	none := oauth.LoadToken("none")
	none.Print()
	broke := oauth.LoadToken("broke")
	broke.Print()
	// Output:
	// {
	//   "access_token": "7e61c8a5e2f99404730c511de6580412e618da35",
	//   "expires": 1520280099,
	//   "expires_in": 3600,
	//   "refresh_token": "0e633c3343a2df84b1526f4c2e6993ff17e05cab",
	//   "scope": "profile.default.read channels.default.read chat.default.read stream.default.read",
	//   "token_type": "Bearer"
	// }
	// null
	// null
}

func ExampleHaveClient() {
	fmt.Println(oauth.HaveClient("some.io"))
	fmt.Println(oauth.HaveClient("nope"))
	// Output:
	// true
	// false
}

func ExampleHaveToken() {
	fmt.Println(oauth.HaveToken("some.io"))
	fmt.Println(oauth.HaveToken("nope"))
	// Output:
	// true
	// false
}
