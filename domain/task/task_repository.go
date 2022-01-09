//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package task

import "context"

type TaskRepository interface {
	Insert(ctx context.Context, task Task) error
	FindByID(ctx context.Context, taskID TaskID) (*Task, error)
	Update(ctx context.Context, task Task) error
}
