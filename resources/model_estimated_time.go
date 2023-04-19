/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EstimatedTime struct {
	Key
	Attributes EstimatedTimeAttributes `json:"attributes"`
}
type EstimatedTimeResponse struct {
	Data     EstimatedTime `json:"data"`
	Included Included      `json:"included"`
}

type EstimatedTimeListResponse struct {
	Data     []EstimatedTime `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustEstimatedTime - returns EstimatedTime from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEstimatedTime(key Key) *EstimatedTime {
	var estimatedTime EstimatedTime
	if c.tryFindEntry(key, &estimatedTime) {
		return &estimatedTime
	}
	return nil
}
