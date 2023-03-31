package data

type RefreshTokens interface {
	New() RefreshTokens

	Create(token RefreshToken) error
	Get() (*RefreshToken, error)
	Delete(token string) error

	FilterByToken(token string) RefreshTokens
}

type RefreshToken struct {
	Token     string `db:"token" structs:"token"`
	OwnerId   int64  `db:"owner_id" structs:"owner_id"`
	ValidDate int64  `db:"valid_date" structs:"valid_date"`
}
