/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserPermissionAttributes struct {
	AccessLevel AccessLevel `json:"access_level"`
	// chat title
	Link string `json:"link"`
	// user id from module
	ModuleId int64 `json:"module_id"`
	// chat title
	Path string `json:"path"`
	// phone from telegram
	Phone *string `json:"phone,omitempty"`
	// submodule access hash to handle submodule with the same title
	SubmoduleAccessHash *string `json:"submodule_access_hash,omitempty"`
	// submodule id to handle submodule with the same title
	SubmoduleId string `json:"submodule_id"`
	// user id from identity
	UserId *int64 `json:"user_id,omitempty"`
	// username from telegram
	Username *string `json:"username,omitempty"`
}
