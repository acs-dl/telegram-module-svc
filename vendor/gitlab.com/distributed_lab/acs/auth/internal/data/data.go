package data

type GenerateTokens struct {
	User              User
	AccessLife        int64
	Secret            string
	PermissionsString string
}

type JwtClaims struct {
	ExpiresAt        int64  `json:"exp"`
	OwnerId          int64  `json:"owner_id"`
	Email            string `json:"email"`
	ModulePermission string `json:"module.permission"`
}

const TokenRegExpStr = "/^([a-zA-Z0-9_=]+)\\.([a-zA-Z0-9_=]+)\\.([a-zA-Z0-9_\\-\\+\\/=]*)/gm"
