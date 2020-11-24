package auth

import (
	"testing"
)

/*
// use the rwxyou (not a real account, not linked to anything) and be
// sure to reset the secret after testing (especially live testing)
func TestConfigSave(t *testing.T) {
	a := new(App)
	a.Name = "rwxyou"
	a.ClientID = "0096494d-4fd6-43bf-ba48-1aa55ebfb04e"
	a.ClientSecret = "6307a405-28a2-4e0d-a647-750f049b40fa"
	a.Endpoint.AuthURL = "https://api.restream.io/login"
	a.Endpoint.TokenURL = "https://api.restream.io/oauth/token"
	c := Config{a}
	dir := t.TempDir()
	file := filepath.Join(dir, "testconfig.json")
	defer os.Setenv("AUTHCONF", os.Getenv("AUTHCONF"))
	os.Setenv("AUTHCONF", file)
	c.Cache()
	buf, _ := ioutil.ReadFile(file)
	fmt.Println(string(buf))
}
*/

func TestConfigClient(t *testing.T) {
	//defer os.Setenv("AUTHCONF", os.Getenv("AUTHCONF"))
	//os.Setenv("AUTHCONF", "testdata/rwxyou/config.json")
	c, err := OpenConfig()
	if err != nil {
		t.Fatal(err)
	}
	_ = c
	//c["rwxyou"].Authorize()
}
