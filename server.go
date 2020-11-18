package oauth

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var localServers map[string]chan interface{}

// LocalServers returns the map containing one entry for all the local
// server goroutines currently running. Each can use the same com
// channel or a different one. Even though the addresses can be
// different all are guaranteed not to have the same port.
func LocalServers() map[string]chan interface{} {
	return localServers
}

// PortTaken returns the address and com channel for a local server that
// is currently using the given port. The argument (this) may be either
// a string address or an integer port number. If a string only the port
// will be used.
func PortTaken(this interface{}) (addr string, com chan interface{}) {
	var port string
	switch v := this.(type) {
	case int:
		port = fmt.Sprintf(":%v", v)
	case string:
		i := strings.Index(v, ":")
		if i < 0 {
			return "", nil
		}
		port = v[i:]
	default:
		return "", nil
	}
	for k, v := range localServers {
		if strings.HasSuffix(k, port) {
			return k, v
		}
	}
	return "", nil
}

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

	if p, _ := PortTaken(addr); p != "" {
		return fmt.Errorf("startserver: already server on port: %v", p)
	}

	handleHome := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hi there, "+d.Name+".\n")
	}

	handleDie := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Oh no, I'm melting ..... ugh.\n")
		com <- "killmenow"
	}

	handleMockLogin := func(w http.ResponseWriter, req *http.Request) {

		// check to see if form data is being submitted
		req.ParseForm()

		// no need to do anything to save the login, just pretend it was
		// successful every time since we are just mocking a login
		redirect_uri := req.FormValue("redirect_uri")
		if redirect_uri != "" {
			http.Redirect(w, req, redirect_uri, 307)
			return
		}

		// by default send a login page
		io.WriteString(w, loginPage)
	}

	handleMockToken := func(w http.ResponseWriter, req *http.Request) {
		// TODO pretend to handle token requests, both initial tokens from
		// code as well as refreshes
		io.WriteString(w, "Would pretend to handle token requests, initial code or refresh\n")
	}

	// encloses the *pointer* to the app data
	handleOauthRedirect := func(w http.ResponseWriter, req *http.Request) {
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
			// TODO check this better
			d.Lock()
			d.Scope = scope
			d.Unlock()
		}

		// TODO consider sending failures to the com as well to decide if
		// fatal and server should be shut down

		// bail unless we have been redirected and have a temp auth code
		if code == "" {
			io.WriteString(w, "Authorization code not found\n")
			return
		}

		if state == "" {
			// TODO also check that matches original login
			io.WriteString(w, "Matching state not found\n")
			return
		}

		io.WriteString(w, "Requesting token using code ... ")
		/*
			err := RequestToken(d)
			if err != nil {
				io.WriteString(w, "FAILED: "+string(err)+"\n")
				return
			}
		*/

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
	http.HandleFunc("/oauth", handleOauthRedirect)
	http.HandleFunc("/login", handleMockLogin)
	http.HandleFunc("/token", handleMockToken)
	// TODO gracefully fail if already running at same address, consider
	// sending message on channel with the addr in case it the com is
	// being shared by multiple servers
	go http.ListenAndServe(addr, nil)

	return nil
}

const loginPage = `
<!doctype html>
<html>
<head>
	<title>Mock Login Page</title>
	<style>

		body {
			font-family: sans-serif;
			background: #2b2825;
			color: #d9cbb5;
		}

		main {
			max-width: 500px;
			margin-left: auto;
			margin-right: auto;
		}

	</style>
</head>
<body>
	<main>
	  <h1>Mock Login Page</h1>

		<p>This mock login page is for simulating integration of a web
		browser as part of the required Oauth2 authorization flow for
		testing and prototyping.</p>

		<p>Just submit if you want to use the default username and
		password (<code>harold</code>, <code>pass1234</code>). These are
		provided so you can test different combinations rather than doing an
		automatic redirect or simply a button with the defaults hard
		coded.</p>

		<form>
			<table>
				<tr>
					<td><label for=user>Username:</label></td>
					<td><input id=user name=user value=harold></td>
				</tr>
				<tr>
					<td><label for=pass>Password:</label></td>
					<td><input id=pass name=pass type=password value=pass1234></td>
				</tr>
				<tr>
					<td></td><td><input type=submit method=post></td>
				</tr>
		  </table>
		</form>

	</main>
</body>
</html>
`
