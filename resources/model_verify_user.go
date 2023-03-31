/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VerifyUser struct {
	// action that must be handled in module, must be \"verify_user\"
	Action string `json:"action"`
	// user's id from identity
	UserId string `json:"user_id"`
	// user's username from gitlab
	Username string `json:"username"`
}
