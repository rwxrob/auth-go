package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// sessions contains all the current authorization sessions in progress
// that the server should be prepared to handle. Keys are usually
// App.State (base32) values but can be anything unique.
var sessions = struct {
	sync.RWMutex
	d map[string]*App
}{d: map[string]*App{}}

// AddSession adds an authorization session for the given app to the
// internal sessions map for the package. SetAuthState() is called on
// the App and the state is used internally as the unique key for the
// session. Safe for concurrency through use of full sync.Mutex.
func AddSession(a *App) {
	defer sessions.Unlock()
	sessions.Lock()
	a.SetAuthState()
	sessions.d[a.AuthState] = a
}

// GetSession returns a session from the internal map cache if found,
// otherwise nil.
func GetSession(state string) *App {
	//	defer sessions.RUnlock()
	//	sessions.RLock()
	if v, has := sessions.d[state]; has {
		return v
	}
	return nil
}

// outputs to browser as soon as written
func write(w http.ResponseWriter, s string) {
	io.WriteString(w, s)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

// HandleRedirect processes all awaiting Oauth2 grant authorization
// sessions checking the states, receiving the code, and then upgrading
// the code to a token writing flushed simple text status messages along
// the way.
func HandleRedirect(w http.ResponseWriter, req *http.Request) {
	var app *App

	write(w, "Parsing data ... ")
	req.ParseForm()
	write(w, "parsed\n")

	write(w, "Checking CSRF state ... ")
	state := req.FormValue("state")
	if state == "" {
		write(w, "state not found\n")
		return
	}
	write(w, "matched\n")

	write(w, "Looking up pending auth sessions ... ")
	app = GetSession(state)
	if app == nil {
		write(w, "not found\n")
		return
	}
	write(w, "found\n")

	write(w, "Looking for authorization code ... ")
	code := req.FormValue("code")
	if code == "" {
		write(w, "not found\n")
		return
	}
	app.SetAuthCode(code)
	write(w, "found\n")

	write(w, "Upgrading to access token ... ")
	ctx := context.Background()
	tok, err := app.Exchange(ctx, app.AuthCode)
	if err != nil {
		write(w, err.Error())
		write(w, "failed\n")
		return
	}
	app.Token = *tok
	delete(sessions.d, app.AuthState) // leave for GC?
	write(w, "authorized\n")

	write(w, "This browser window can now be closed.")
}

// StartLocalServer start the main HTTP server locally to receive
// redirects from Authorize. The same server is used for everything that
// requires one in this package so care has been taken to ensure
// requests are handled such that they are safe for concurrency.  The
// hard coded to the address localhost:8080 (and not :8080 which would
// expose the server to other external interfaces).
func StartLocalServer() error {
	http.HandleFunc("/redirected", HandleRedirect)
	addr := "localhost:8080"
	fmt.Printf("Starting server (%v)\n", addr)
	go http.ListenAndServe(addr, nil)
	go startSessionGC()
	return nil
}

func startSessionGC() {
	// TODO add a garbage collection for Sessions for stuff that becomes
	// defunct or gets resubmitted, else memory leaks, even though this
	// isn't intended for a high reliability scope and more for terminal
	// applications.
}
