/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Roles struct {
	Key
	Attributes RolesAttributes `json:"attributes"`
}
type RolesResponse struct {
	Data     Roles    `json:"data"`
	Included Included `json:"included"`
}

type RolesListResponse struct {
	Data     []Roles  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustRoles - returns Roles from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRoles(key Key) *Roles {
	var roles Roles
	if c.tryFindEntry(key, &roles) {
		return &roles
	}
	return nil
}
