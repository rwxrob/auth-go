/*
Package auth contains mostly Oauth2-specific convenience functions and structures to pick up where the official package leaves of. For example, caching of tokens and other related data in a way that is localized to the current user and easily retrievable by scripts and other Go code running with the permissions of that user.

The auth command is included under cmd. It is designed to be used by shell scripts to facilitate use of a common authentication cache for the current user.
*/
package auth
