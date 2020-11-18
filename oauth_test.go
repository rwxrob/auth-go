package oauth

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestPortTaken(t *testing.T) {
	com := make(chan interface{})
	localServers = map[string]chan interface{}{
		"localhost:8080": com,
		"localhost:8081": com,
		":8082":          com,
	}

	check := func(this interface{}, addr string) bool {
		gotaddr, gotcom := PortTaken(this)
		return gotaddr == addr && gotcom != nil
	}

	tests := [][]string{
		{"localhost:8080", "localhost:8080"},
		{":8080", "localhost:8080"},
		{"localhost:8081", "localhost:8081"},
		{":8081", "localhost:8081"},
		{":8082", ":8082"},
		{"localhost:8082", ":8082"},
		{"another:8082", ":8082"},
	}
	for _, test := range tests {
		t.Logf("%v -> %v", test[0], test[1])
		if !check(test[0], test[1]) {
			t.Errorf("failed: %v", test[0])
		}
	}

}

func TestUpdate(t *testing.T) {
	var err error
	var d *AppData

	defer os.Setenv("OAUTHDIR", os.Getenv("OAUTHDIR"))
	dir := t.TempDir()
	os.Setenv("OAUTHDIR", dir)

	// create a pretend server to check "port in use"
	tcom := make(chan interface{})
	localServers = map[string]chan interface{}{
		"some:8090": tcom,
	}

	write := func(this string) {
		ioutil.WriteFile(
			filepath.Join(dir, "appdata.json"),
			[]byte(this), 0600,
		)
	}

	/*
		get := func() string {
			buf, _ := ioutil.ReadFile(filepath.Join(dir, "appdata.json"))
			return string(buf)
		}
	*/

	// no cache by that name at all, app data must already exist
	err = Update("empty")
	t.Log(err)
	if err == nil {
		t.Error("should have error if failed to load app data")
		return
	}

	// cached but missing the client_id
	write(`{}`)
	err = Update("appdata")
	t.Log(err)
	if err == nil {
		t.Error("should have failed from no client_id")
		return
	}
	err = nil

	// cached but missing the client_secret
	write(`{"name":"appdata","client_id":"someid"}`)
	err = Update("appdata")
	t.Log(err)
	if err == nil {
		t.Error("should have failed from no client_id")
		return
	}
	err = nil

	// already a server on the port of the address in appdata
	write(`{"name":"appdata","client_id":"someid", "client_secret":"somesecret","redirect_uri":":8090"}`)
	d = Load("appdata")
	err = Update(d)
	t.Log(err)
	if err == nil {
		t.Error("should have failed with conflicting port")
		return
	}
	err = nil

	// startup the mock server to simulate login authorization flow
	mockcom := make(chan interface{})
	StartServer(d, "localhost:8091", mockcom)
	<-mockcom
	// cached with expires_in but no expires should set expires and ignore
	write(`{"name":"appdata","client_id":"someid","client_secret":"somesecret","expires_in":3600}`)
	d = Load("appdata")
	err = Update(d)
	if err != nil {
		t.Error("should not have failed just for no expires (which is updated)")
		return
	}
	if Load("appdata").Expires == 0 {
		t.Error("should have updated expires")
		return
	}
	err = nil

	/*
		// cached with more than 10 minutes left before expires should
		// just ignore and return
		write(`{"name":"appdata","client_id":"someid","client_secret":"somesecret","expires_in":3600}`)
		d = Load("appdata")
		d.SetExpires()
		d.Expires += 10000

		err = Update(d)
		if err != nil {
			t.Error("should have just ignored >10 minute expires")
			return
		}
		err = nil
	*/

	// cached with refresh token should get a refresh for a new token

	// a failed refresh token request (for any reason) should trigger auth

	// everything else should trigger a fully interactive auth

	<-mockcom
	return

}

/*
func TestOpenLocalBrowser(t *testing.T) {
	OpenLocalBrowser("https://robs.io")
}
*/
