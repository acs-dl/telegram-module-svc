/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Submodules struct {
	Key
	Attributes SubmodulesAttributes `json:"attributes"`
}
type SubmodulesResponse struct {
	Data     Submodules `json:"data"`
	Included Included   `json:"included"`
}

type SubmodulesListResponse struct {
	Data     []Submodules `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustSubmodules - returns Submodules from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSubmodules(key Key) *Submodules {
	var submodules Submodules
	if c.tryFindEntry(key, &submodules) {
		return &submodules
	}
	return nil
}
