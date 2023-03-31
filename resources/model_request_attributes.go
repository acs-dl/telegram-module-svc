/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type RequestAttributes struct {
	// Module to grant permission
	Module string `json:"module"`
	// Already built payload to grant permission <br><br> -> \"add_user\" = action to add user in chat in telegram<br> -> \"verify_user\" = action to verify user in telegram module (connect user id from identity with telegram info)<br> -> \"update_user\" = action to update user access level in chat in telegram<br> -> \"get_users\" = action to get users with their permissions from chats in telegram<br> -> \"delete_user\" = action to delete user from module (from all links)<br> -> \"remove_user\" = action to remove user from chat in telegram<br>
	Payload json.RawMessage `json:"payload"`
}
