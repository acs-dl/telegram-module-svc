package lorem

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(source string) string {
	digest := sha256.Sum256([]byte(source))
	return hex.EncodeToString(digest[:])
}
