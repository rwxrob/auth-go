package oauth

import (
	"os"
	"testing"
)

func TestToJSON(t *testing.T) {
	null := toJSON(nil)
	if null != "null" {
		t.Error("toJSON: expected null")
	}
}

func TestPath(t *testing.T) {
	p := path("some")
	if p != "" {
		t.Error("path: should be empty if not exist")
	}
	prev := os.Getenv("OAUTHDIR")
	defer os.Setenv("OAUTHDIR", prev)
	os.Setenv("OAUTHDIR", "")
	p = path("some.io")
	if p != "" {
		t.Error("path: should be empty if no OAUTHDIR env set")
	}

}
