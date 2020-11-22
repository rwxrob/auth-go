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
	defer sessions.RUnlock()
	sessions.RLock()
	if v, has := sessions.d[state]; has {
		return v
	}
	return nil
}

// HandleRedirect processes all awaiting Oauth2 grant authorization
// sessions checking the states, receiving the code, and then upgrading
// the code to a token.
func HandleRedirect(w http.ResponseWriter, req *http.Request) {
	var app *App

	io.WriteString(w, "Parsing data ... ")
	req.ParseForm()
	io.WriteString(w, "parsed\n")

	io.WriteString(w, "Checking CSRF state ... ")
	state := req.FormValue("state")
	if state == "" {
		io.WriteString(w, "state not found\n")
		return
	}
	io.WriteString(w, "matched\n")

	io.WriteString(w, "Looking up pending auth sessions ... ")
	app = GetSession(state)
	if app == nil {
		io.WriteString(w, "not found\n")
		return
	}
	io.WriteString(w, "found\n")

	// concurrent writes ahead
	app.Lock()
	defer app.Unlock()

	io.WriteString(w, "Looking for authorization code ... ")
	code := req.FormValue("code")
	if code == "" {
		io.WriteString(w, "not found\n")
		return
	}
	app.SetAuthCode(code)
	io.WriteString(w, "found\n")

	io.WriteString(w, "Upgrading to access token ... ")
	ctx := context.Background()
	tok, err := app.Exchange(ctx, app.AuthCode)
	if err != nil {
		io.WriteString(w, err.Error())
		io.WriteString(w, "failed\n")
		return
	}
	app.Token = *tok
	delete(sessions.d, app.AuthState)
	io.WriteString(w, "authorized\n")
}

// StartLocalServer start the main HTTP server locally to receive
// redirects from Authorize. The same server is used for everything that
// requires one in this package so care has been taken to ensure
// requests are handled such that they are safe for concurrency.  By
// default the server is started on the address localhost:8080 (and not
// :8080 which would expose the server to other external interfaces) but
// can be overriden with the AUTHHTTP environment variable. Also see
// HandleRedirect.
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
