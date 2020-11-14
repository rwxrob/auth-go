package auth

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Get simply fetches a value from its file within the AUTHDIR. Only the
// first line is ever read allowing comments and such to be written on
// subsequent lines.
func Get(path string) string {
	fpath := filepath.Join(os.Getenv("AUTHDIR"), path)
	f, _ := os.Open(fpath)
	defer f.Close()
	v, _, _ := bufio.NewReader(f).ReadLine()
	return strings.Trim(string(v), " \n")
}
