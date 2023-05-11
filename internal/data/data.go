package data

import (
	"time"
)

const (
	ModuleName        = "telegram"
	UnverifiedService = "unverified-svc"
	IdentityService   = "identity"
)

const InviteMessageTemplate = `Hello, <b>{{.Name}}</b> !

We have tried to add you in <i>{{.GroupName}}</i> group, but can't.

Here is your invite link: <a href={{.InviteLink}}>CLICK HERE</a>

Note that link is only for <b>you</b> and valid for <b>1 hour</b>`

type ModuleRequest struct {
	ID            string    `db:"id" structs:"id"`
	UserID        int64     `db:"user_id" structs:"user_id"`
	Module        string    `db:"module" structs:"module"`
	Payload       string    `db:"payload" structs:"payload"`
	CreatedAt     time.Time `db:"created_at" structs:"created_at"`
	RequestStatus string    `db:"request_status" structs:"request_status"`
	Error         string    `db:"error" structs:"error"`
}

type ModulePayload struct {
	RequestId string `json:"request_id"`
	UserId    string `json:"user_id"`
	Action    string `json:"action"`

	//other fields that are required for module
	Link        string   `json:"link"`
	Links       []string `json:"links"`
	Username    *string  `json:"username"`
	Phone       *string  `json:"phone"`
	AccessLevel string   `json:"access_level"`
}

type UnverifiedPayload struct {
	Action string           `json:"action"`
	Users  []UnverifiedUser `json:"users"`
}

var Roles = map[string]string{
	Admin:  "Admin",
	Member: "Member",
	Owner:  "Owner",

	Self:   "Self",
	Left:   "Left",
	Banned: "Banned",
}

type MessageInfo struct {
	Data            map[string]interface{}
	MessageTemplate string
	User            User
}
