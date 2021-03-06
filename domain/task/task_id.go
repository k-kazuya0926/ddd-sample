package task

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type TaskID struct {
	id ulid.ULID
}

// ドメイン層からしか呼び出せないよう、パッケージプライベートとしている
func newTaskID() TaskID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return TaskID{id: ulid.MustNew(ulid.Timestamp(t), entropy)}
}

func (t TaskID) String() string {
	return t.id.String()
}

func ParseTaskID(id string) (TaskID, error) {
	ulid, err := ulid.Parse(id)
	if err != nil {
		return TaskID{}, err
	}
	return TaskID{id: ulid}, nil
}
