package auth

import "testing"

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

func TestRedirectHost(t *testing.T) {
	a := new(App)
	a.RedirectURL = "http://localhost:8080/oath"
	host := a.RedirectHost()
	t.Log(host)
	if host != "localhost:8080" {
		t.Fail()
	}
}

/*
func TestExpired(t *testing.T) {
	a := new(App)
	a.Expiry, _ = time.Parse(time.RFC3339Nano, "2020-11-21T07:43:14.178172151-05:00")
	fmt.Println(a.Expiry.Before(time.Now()))
}
*/
