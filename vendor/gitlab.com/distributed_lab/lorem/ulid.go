package lorem

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// ULID see https://github.com/ulid/spec for format details
// NOTE: Not battle-tested yet, use with caution.
func ULID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader).String()
}
