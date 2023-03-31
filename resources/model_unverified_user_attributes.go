/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type UnverifiedUserAttributes struct {
	// timestamp without timezone when user was created
	CreatedAt time.Time `json:"created_at"`
	// email from telegram
	Email *string `json:"email,omitempty"`
	// module name
	Module string `json:"module"`
	// user id from module
	ModuleId int64 `json:"module_id"`
	// name from telegram
	Name *string `json:"name,omitempty"`
	// phone from module
	Phone *string `json:"phone,omitempty"`
	// username from module
	Username *string `json:"username,omitempty"`
}
