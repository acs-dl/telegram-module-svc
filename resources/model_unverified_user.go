/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UnverifiedUser struct {
	Key
	Attributes UnverifiedUserAttributes `json:"attributes"`
}
type UnverifiedUserResponse struct {
	Data     UnverifiedUser `json:"data"`
	Included Included       `json:"included"`
}

type UnverifiedUserListResponse struct {
	Data     []UnverifiedUser `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustUnverifiedUser - returns UnverifiedUser from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUnverifiedUser(key Key) *UnverifiedUser {
	var unverifiedUser UnverifiedUser
	if c.tryFindEntry(key, &unverifiedUser) {
		return &unverifiedUser
	}
	return nil
}
