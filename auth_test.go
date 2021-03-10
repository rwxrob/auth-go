package auth

import (
	"os"
	"testing"
)

func TestValid(t *testing.T) {
	defer os.Setenv("AUTHCONF", os.Getenv("AUTHCONF"))
	os.Setenv("AUTHCONF", "testdata/auth.json")
	// TODO not yet expired
}
