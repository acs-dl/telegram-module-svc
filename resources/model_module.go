/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Module struct {
	Key
	Attributes ModuleAttributes `json:"attributes"`
}
type ModuleResponse struct {
	Data     Module   `json:"data"`
	Included Included `json:"included"`
}

type ModuleListResponse struct {
	Data     []Module `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustModule - returns Module from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustModule(key Key) *Module {
	var module Module
	if c.tryFindEntry(key, &module) {
		return &module
	}
	return nil
}
