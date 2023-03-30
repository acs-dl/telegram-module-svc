package data

import (
	"time"
)

const (
	ModuleName        = "telegram"
	UnverifiedService = "telegram-module"
	IdentityService   = "identity"
)

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
	Link        string  `json:"link"`
	Username    *string `json:"username"`
	Phone       *string `json:"phone"`
	AccessLevel string  `json:"access_level"`
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
