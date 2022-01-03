package task

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type TaskID struct {
	id ulid.ULID
}

func NewTaskID() TaskID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return TaskID{id: ulid.MustNew(ulid.Timestamp(t), entropy)}
}
