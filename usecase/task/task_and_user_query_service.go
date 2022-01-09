package task

import (
	"ddd-sample/domain/shared/page"
	"ddd-sample/domain/task"
	"ddd-sample/domain/user"
)

type TaskAndUserDTO struct {
	TaskID   task.TaskID
	TaskName task.TaskName
	UserName user.UserName
}

type TaskAndUserDTOPage struct {
	// 取得したエンティティ
	TaskAndUserDTOs []TaskAndUserDTO

	// ページング情報
	Paging page.Paging
}

type TaskAndUserQueryService interface {
	FetchList(pagingCondition page.PagingCondition) (TaskAndUserDTOPage, error)
}
