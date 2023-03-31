/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserInfo struct {
	Key
	Attributes UserInfoAttributes `json:"attributes"`
}
type UserInfoResponse struct {
	Data     UserInfo `json:"data"`
	Included Included `json:"included"`
}

type UserInfoListResponse struct {
	Data     []UserInfo `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustUserInfo - returns UserInfo from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUserInfo(key Key) *UserInfo {
	var userInfo UserInfo
	if c.tryFindEntry(key, &userInfo) {
		return &userInfo
	}
	return nil
}
