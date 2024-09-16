package libs

import (
	"time"

	"math/rand"

	"github.com/Chankuri2049/Sand/types"
	"github.com/oklog/ulid/v2"
)

var entropy ulid.MonotonicReader

func init() {
	entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
}

func setULID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}

func SetULID() types.ID {
	return types.ID(setULID())
}
