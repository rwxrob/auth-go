package oauth_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/rwxrob/oauth"
)

func ExampleSave() {
	app := new(oauth.AppData)
	app.Name = "random"
	app.ClientId = "random_client_id"
	app.ClientSecret = "random_client_secret"
	err := app.Save("testdata/save.json")
	if err != nil {
		fmt.Println(err)
	}
	data, _ := ioutil.ReadFile("testdata/save.json")
	fmt.Println(string(data))
	// Output:
	// {
	//   "app_name": "random",
	//   "client_id": "random_client_id",
	//   "client_secret": "random_client_secret"
	// }
}

func ExampleCache() {
	defer os.Setenv("OAUTHDIR", os.Getenv("OAUTHDIR"))
	os.Setenv("OAUTHDIR", "testdata")
	app := new(oauth.AppData)
	app.Name = "random"
	app.ClientId = "random_client_id"
	app.ClientSecret = "random_client_secret"
	err := oauth.Cache(app)
	if err != nil {
		fmt.Println(err)
	}
	oauth.Load("random").Print()
	// Output:
	// {
	//   "app_name": "random",
	//   "client_id": "random_client_id",
	//   "client_secret": "random_client_secret"
	// }
}

func ExampleLoad() {
	os.Setenv("OAUTHDIR", "testdata")
	client := oauth.Load("some.io")
	client.Print()
	none := oauth.Load("none")
	none.Print()
	broke := oauth.Load("broke")
	broke.Print()
	// Output:
	// {
	//   "client_id": "some_client_id",
	//   "client_secret": "some_client_secret",
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

func ExampleHave() {
	fmt.Println(oauth.Have("some.io"))
	fmt.Println(oauth.Have("nope"))
	// Output:
	// true
	// false
}