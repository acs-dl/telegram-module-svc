/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Chat struct {
	Key
	Attributes ChatAttributes `json:"attributes"`
}
type ChatResponse struct {
	Data     Chat     `json:"data"`
	Included Included `json:"included"`
}

type ChatListResponse struct {
	Data     []Chat   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustChat - returns Chat from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustChat(key Key) *Chat {
	var chat Chat
	if c.tryFindEntry(key, &chat) {
		return &chat
	}
	return nil
}
