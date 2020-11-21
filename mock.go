package auth

// StartMockEndpoint starts a new local HTTP server to serve as a mock
// endpoint for debugging Oauth2 authentication flows. The address
// provided will be substituted for the scheme, domain, and port portion
// of the Endpoint.AuthURL and Endpoint.TokenURL from the configuration
// data passed to it. The remaining part (route) of the URLs will be
// used for the server routes.
func StartMockEndpoint(addr string, d *Data, com chan interface{}) error {
	// TODO make sure there aren't alread Go packages that do this
	return nil
}
