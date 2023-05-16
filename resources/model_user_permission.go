/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserPermission struct {
	Key
	Attributes UserPermissionAttributes `json:"attributes"`
}
type UserPermissionResponse struct {
	Data     UserPermission `json:"data"`
	Included Included       `json:"included"`
}

type UserPermissionListResponse struct {
	Data     []UserPermission `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustUserPermission - returns UserPermission from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUserPermission(key Key) *UserPermission {
	var userPermission UserPermission
	if c.tryFindEntry(key, &userPermission) {
		return &userPermission
	}
	return nil
}
