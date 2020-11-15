package oauth

import "net/http"

type Server struct {
	http.Server
	Port       int              // local port (must be available)
	Chan       chan interface{} // unbuffered two way comm channel
	Route      string           // default: oauth
	Redirected string           // default: hook (oauth/redirected)
	Login      string           // default: login (oauth/login) (mock)
}

// NewServer returns a local HTTP server (ex: http://localhost:8080)
// that can be used for many things:
//
// * Local redirect_uri during authentication flow
// * Mock Oauth2 complient authentication server for testing flows
//
func NewServer() *Server {
	s := new(Server)
	s.Chan = make(chan interface{})
	s.Route = "oauth"
	s.Redirected = "redirected"
	s.Login = "login"
	return s
}

// Once started the fields of the server struct will be ignored if
// changed.
