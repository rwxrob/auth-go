package auth

// Collection contains one or more Data structs such that they can be
// cached and managed as a single secure file. This is particularly
// useful when dealing with single-passphrase encryption approaches to
// cache storage.
type Collection []*Data

// SaveWithTouchAuth securely saves the collection to disk with
// read/write permission only for the current user with a physical
// secure device key (like YubiKey) that they can touch. Any existing
// file at the path location will be overwritten (permissions
// permitting). Every call to this method will require a physical
// interaction with the device connected to the host computer.
func (c *Collection) SaveWithTouchAuth(path string) error {
	// TODO
	return nil
}
