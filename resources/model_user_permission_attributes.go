/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserPermissionAttributes struct {
	AccessLevel AccessLevel `json:"access_level"`
	// link level in tree
	Level int64 `json:"level"`
	// chat title
	Link string `json:"link"`
	// user id from module
	ModuleId int64 `json:"module_id"`
	// phone from telegram
	Phone string `json:"phone"`
	// user id from identity
	UserId *int64 `json:"user_id,omitempty"`
	// username from telegram
	Username string `json:"username"`
}
