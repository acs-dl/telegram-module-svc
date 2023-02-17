/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AddUser struct {
	// user's id from identity
	AccessLevel int `json:"access_level"`
	// action that must be handled in module, must be \"add_user\"
	Action string `json:"action"`
	// link where module has to add user
	Link string `json:"link"`
	// user's id from identity
	UserId string `json:"user_id"`
	// user's username from gitlab
	Username string `json:"username"`
}
