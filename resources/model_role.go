/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Role struct {
	Key
	Attributes RoleAttributes `json:"attributes"`
}
type RoleResponse struct {
	Data     Role     `json:"data"`
	Included Included `json:"included"`
}

type RoleListResponse struct {
	Data     []Role   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustRole - returns Role from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRole(key Key) *Role {
	var role Role
	if c.tryFindEntry(key, &role) {
		return &role
	}
	return nil
}
