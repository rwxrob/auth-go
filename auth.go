package auth

import "os/exec"

// OpenLocalBrowser opens the users local graphic browser (with
// xdg-open)
func OpenLocalBrowser(url string) {
	// TODO: detect other platforms
	exec.Command("xdg-open", url).Start()
}
