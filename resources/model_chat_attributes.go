/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChatAttributes struct {
	// telegram chat members amount
	MembersAmount int64 `json:"members_amount"`
	// link to the chat photo
	Photo *string `json:"photo,omitempty"`
	// telegram chat access hash
	SubmoduleAccessHash *string `json:"submodule_access_hash,omitempty"`
	// telegram chat id
	SubmoduleId string `json:"submodule_id"`
	// telegram chat title
	Title string `json:"title"`
}
