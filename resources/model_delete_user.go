/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DeleteUser struct {
	// action that must be handled in module, must be \"delete_user\"
	Action string `json:"action"`
	// user's username from gitlab
	Username string `json:"username"`
}
