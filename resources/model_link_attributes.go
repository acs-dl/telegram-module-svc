/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type LinkAttributes struct {
	Chats []Chat `json:"chats"`
	// indicates whether link exists
	IsExists bool `json:"is_exists"`
	// link to repository or group
	Link string `json:"link"`
}
