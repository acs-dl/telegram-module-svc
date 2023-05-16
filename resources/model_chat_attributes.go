/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChatAttributes struct {
	// telegram chat access hash
	AccessHash *int64 `json:"access_hash,omitempty"`
	// telegram chat id
	Id int64 `json:"id"`
	// telegram chat members amount
	MembersAmount int64 `json:"members_amount"`
	// link to the chat photo
	Photo *string `json:"photo,omitempty"`
	// telegram chat title
	Title string `json:"title"`
}
