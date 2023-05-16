/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Inputs struct {
	Key
	Attributes InputsAttributes `json:"attributes"`
}
type InputsResponse struct {
	Data     Inputs   `json:"data"`
	Included Included `json:"included"`
}

type InputsListResponse struct {
	Data     []Inputs `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustInputs - returns Inputs from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustInputs(key Key) *Inputs {
	var inputs Inputs
	if c.tryFindEntry(key, &inputs) {
		return &inputs
	}
	return nil
}
