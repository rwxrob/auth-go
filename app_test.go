package auth

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseRedirectURL(t *testing.T) {
	a := new(App)
	a.RedirectURL = "http://localhost:8080/oath"
	uri, _ := a.ParseRedirectURL()
	uri.Path = ""
	t.Log(uri)
	if uri.String() != "http://localhost:8080" {
		t.Fail()
	}
}

func TestAppString(t *testing.T) {
	a := new(App)
	t.Log(a.String())
}

func TestAppPrint(t *testing.T) {
	a := new(App)
	a.Print()
}

func TestAppJSON(t *testing.T) {
	a := new(App)
	t.Log(string(a.JSON()))
}

func TestAppSetAuthState(t *testing.T) {
	a := new(App)
	a.SetAuthState()
	if a.AuthState == "" {
		t.Error("AuthState not set.")
		return
	}
	t.Log(a)
}

func TestAppSetAuthCode(t *testing.T) {
	a := new(App)
	a.SetAuthCode("somecode")
	if a.AuthCode != "somecode" {
		t.Error("AuthCode not set.")
		return
	}
	t.Log(a)
}

func TestAppRefreshNow(t *testing.T) {
	a := new(App)
	err := a.RefreshNow()
	if err == nil {
		t.Fail()
	}
	// TODO eventually mock a server to test underlying oauth calls
}

func TestAppSave(t *testing.T) {
	a := new(App)
	file, _ := ioutil.TempFile(os.TempDir(), "")
	defer os.Remove(file.Name())
	t.Log(file.Name())
	a.Save(file.Name())
	// TODO load the tmpfile to compare
}

func TestRedirectHost(t *testing.T) {
	a := new(App)
	a.RedirectURL = "http://localhost:8080/oath"
	host := a.RedirectHost()
	t.Log(host)
	if host != "localhost:8080" {
		t.Fail()
	}
}

func TestAppLoad(t *testing.T) {
	a := new(App)
	err := a.Load("idontexist")
	if err == nil {
		t.Error("testappload: failed to return error for non-existent file")
		return
	}
	a.Load("testdata/humm.json")
	t.Log(a)
}

/*
func TestExpired(t *testing.T) {
	a := new(App)
	a.Expiry, _ = time.Parse(time.RFC3339Nano, "2020-11-21T07:43:14.178172151-05:00")
	fmt.Println(a.Expiry.Before(time.Now()))
}
*/
