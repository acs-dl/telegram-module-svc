package data

type RefreshTokens interface {
	New() RefreshTokens

	Create(token RefreshToken) error
	Get() (*RefreshToken, error)
	Select() ([]RefreshToken, error)
	Delete() error

	FilterByTokens(tokens ...string) RefreshTokens
	FilterByLowerValidTill(expiresAtUnix int64) RefreshTokens
}

type RefreshToken struct {
	Token     string `db:"token" structs:"token"`
	OwnerId   int64  `db:"owner_id" structs:"owner_id"`
	ValidTill int64  `db:"valid_till" structs:"valid_till"`
}
