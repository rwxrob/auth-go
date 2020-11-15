package oauth

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// OpenLocalBrowser opens the users local graphic browser (with
// xdg-open) TODO: detect other platforms
func OpenLocalBrowser(url string) {
	exec.Command("xdg-open", url).Start()
}

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
	dir := os.Getenv("OAUTHDIR")
	if dir == "" {
		log.Print("loadclient: OAUTHDIR not set")
		return ""
	}
	these := []string{dir}
	these = append(these, this...)
	return filepath.Join(these...)
}
