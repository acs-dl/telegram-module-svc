/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Link struct {
	Key
	Attributes LinkAttributes `json:"attributes"`
}
type LinkResponse struct {
	Data     Link     `json:"data"`
	Included Included `json:"included"`
}

type LinkListResponse struct {
	Data     []Link   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustLink - returns Link from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLink(key Key) *Link {
	var link Link
	if c.tryFindEntry(key, &link) {
		return &link
	}
	return nil
}
