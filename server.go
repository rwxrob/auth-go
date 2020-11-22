package auth

import (
	"context"
	"io"
	"net/http"
)

// Sessions contains all the current authorization sessions in progress
// that the server should be prepared to handle. Keys are usually
// App.State (base32) values but can be anything unique.
var Sessions map[string]*App

// ServerChan is for sending and receiving messages to the single,
// internal HTTP server that handles all local requests such as
// redirects triggered from Authorize().
var ServerChan chan interface{}

// HandleRedirect processes all awaiting Oauth2 grant authorization
// sessions checking the states, receiving the code, and then upgrading
// the code to a token.
func HandleRedirect(w http.ResponseWriter, req *http.Request) {
	var app *App

	io.WriteString(w, "Parsing data ... ")
	req.ParseForm()
	io.WriteString(w, "parsed]n")

	io.WriteString(w, "Checking CSRF state ... ")
	state := req.FormValue("state")
	if state == "" {
		io.WriteString(w, "state not found\n")
		return
	}
	io.WriteString(w, "matched\n")

	io.WriteString(w, "Looking up pending auth sessions ... ")
	app, has := Sessions[state]
	if !has {
		io.WriteString(w, "not found\n")
	}
	io.WriteString(w, "found\n")

	// ensure safety for concurrency
	app.Lock()
	defer app.Unlock()

	io.WriteString(w, "Looking for authorization code ... ")
	app.AuthCode = req.FormValue("code")
	if app.AuthCode == "" {
		io.WriteString(w, "not found\n")
		return
	}
	io.WriteString(w, "found\n")

	io.WriteString(w, "Upgrading to access token ... ")
	ctx := context.Background()
	tok, err := app.Exchange(ctx, app.AuthCode)
	if err != nil {
		io.WriteString(w, "failed\n")
		return
	}
	app.Token = *tok
	io.WriteString(w, "authorized\n")
}

// StartLocalServer start the main HTTP server locally to receive
// redirects from Authorize. The same server is used for everything that
// requires one in this package so care has been taken to ensure
// requests are handled such that they are safe for concurrency. All
// communication to and from the server comes either through the
// ServerChan channel. By default the server is started on the address
// localhost:8080 (and not :8080 which would expose the server to other
// external interfaces) but can be overriden with the AUTHHTTP
// environment variable. Also see HandleRedirect.
func StartLocalServer() error {
	http.HandleFunc("/redirected", HandleRedirect)
	addr := "localhost:8080"
	http.ListenAndServe(addr, nil)
	startSessionGC()
FOR:
	for {
		msg := <-ServerChan
		switch v := msg.(type) {
		case string:
			switch v {
			case "die":
				break FOR
			}
		case *App:
			Sessions[v.AuthState] = v
		}
	}
	return nil
}

func startSessionGC() {
	// TODO add a garbage collection for Sessions for stuff that becomes
	// defunct or gets resubmitted, else memory leaks, even though this
	// isn't intended for a high reliability scope and more for terminal
	// applications.
}
