package user

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// 値オブジェクト。ulid.ULIDへの依存箇所を少なくするために導入した
type UserID struct {
	id ulid.ULID
}

// ドメイン層からしか呼び出せないよう、パッケージプライベートとしている
func newUserID() UserID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return UserID{id: ulid.MustNew(ulid.Timestamp(t), entropy)}
}

func ParseUserID(id string) (UserID, error) {
	ulid, err := ulid.Parse(id)
	if err != nil {
		return UserID{}, err
	}
	return UserID{id: ulid}, nil
}
