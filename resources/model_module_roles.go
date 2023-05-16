/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ModuleRoles struct {
	Key
	Attributes ModuleRolesAttributes `json:"attributes"`
}
type ModuleRolesResponse struct {
	Data     ModuleRoles `json:"data"`
	Included Included    `json:"included"`
}

type ModuleRolesListResponse struct {
	Data     []ModuleRoles `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustModuleRoles - returns ModuleRoles from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustModuleRoles(key Key) *ModuleRoles {
	var moduleRoles ModuleRoles
	if c.tryFindEntry(key, &moduleRoles) {
		return &moduleRoles
	}
	return nil
}
