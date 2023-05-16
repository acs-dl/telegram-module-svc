/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DeleteUser struct {
	// action that must be handled in module, must be \"delete_user\"
	Action string `json:"action"`
	// phone from telegram
	Phone *string `json:"phone,omitempty"`
	// user's username from telegram
	Username *string `json:"username,omitempty"`
}
