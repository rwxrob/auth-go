package oauth

import (
	"fmt"
	"io"
	"net/http"
)

// StartServer starts a local HTTP server for an app that uses Oauth.
// This can be used for many things:
//
// * Local redirect_uri during authentication flow
// * Mock Oauth2 complient authentication server for testing flows
//
// The server will keep the application that started it running until
// a "stop" is sent to the returned channel (or the program terminates).
//
// Even though there is only a single local server and the AppData
// associated with the server can only be for one app at a time, new
// AppData can be passed to the com channel to change it at any time.
// This will change the state of the server to work for the new app
// without restarting it.
//
func StartServer(d *AppData, addr string, com chan interface{}) error {

	// assert good arguments
	switch {
	case d == nil:
		return fmt.Errorf("startserver: pointer to AppData is required")
	case addr == "":
		return fmt.Errorf("startserver: server address (ex: localhost:8080) is required")
	case com == nil:
		return fmt.Errorf("startserver: com channel cannot be nil")
	}

	handleHome := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hi there, "+d.Name+".\n")
	}

	handleDie := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Oh no, I'm melting ..... ugh.\n")
		com <- "killmenow"
	}

	handleMock := func(w http.ResponseWriter, req *http.Request) {
		// TODO
		io.WriteString(w, "Would pretend to be an Oauth web site.\n")
	}

	// encloses the *pointer* to the app data
	handleOauth := func(w http.ResponseWriter, req *http.Request) {
		var code, state, scope string

		// get as much stuff as we can expect
		if v, has := req.Form["code"]; has {
			code = v[0]
		}
		if v, has := req.Form["state"]; has {
			state = v[0]
		}
		if v, has := req.Form["scope"]; has {
			scope = v[0]
		}

		// TODO consider sending failures to the com as well to decide if
		// fatal and server should be shut down

		// bail unless we have been redirected and have a temp auth code
		if code == "" {
			io.WriteString(w, "Authorization code not found\n")
			return
		}

		io.WriteString(w, "Requesting token using code ... ")

		err := RequestToken(d)
		if err != nil {
			io.WriteString(w, "FAILED: "+string(err)+"\n")
			return
		}

		Cache(d)
		io.WriteString(w, "cached.\n")
		// TODO send a com that we cached something

	}

	// The oauth flow requires the same redirect_uri to be used for ever
	// step in the flow even if the redirect_uri is not involved. You
	// cannot start the authorization process and send one
	// localhost:8080/oauth and then use localhost:8080/token for
	// upgrading the code to a token later.
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/die", handleDie)
	http.HandleFunc("/oauth", handleOauth)
	http.HandleFunc("/mock", handleMock)
	go http.ListenAndServe(addr, nil)

	return nil
}
