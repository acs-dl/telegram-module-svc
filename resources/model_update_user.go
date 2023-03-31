/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateUser struct {
	// action that must be handled in module, must be \"update_user\"
	Action string `json:"action"`
	// link where module has to add user
	Link string `json:"link"`
	// phone from telegram
	Phone *string `json:"phone,omitempty"`
	// user's username from telegram
	Username *string `json:"username,omitempty"`
}
