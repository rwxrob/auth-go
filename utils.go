package auth

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func toJSON(thing interface{}) string {
	if thing == nil {
		return "null"
	}
	byt, _ := json.MarshalIndent(thing, "", "  ")
	return string(byt)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func path(this ...string) string {
	dir := os.Getenv("AUTHCACHE")
	if dir == "" {
		return ""
	}
	these := []string{dir}
	these = append(these, this...)
	return filepath.Join(these...)
}
