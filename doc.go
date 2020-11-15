/*
Package oauth contains helper functions and structures designed to make
using Oauth2 easier from the command line. Data is cached in the
OAUTHDIR, which must be set in order to use most of the package. Data is
cached in plain text with restrictive permissions (as is common for ssh
files and such as well). Eventually, a global password for the cache
will be implemented so that nothing is stored unencrypted on disk.

The package comes with an oauth utility command as well that can be
integrated into shell scripts easily for prototyping and the like before
the decision to port to a more substantial language later.

Design Considerations

* Functional paradigm whenever possible

See Also

* https://tools.ietf.org/html/rfc6749

*/
package oauth
