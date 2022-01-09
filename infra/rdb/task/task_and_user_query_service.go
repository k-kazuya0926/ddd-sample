package task

import (
	"ddd-sample/domain/shared/page"
	"ddd-sample/usecase/task"
	usecase_task "ddd-sample/usecase/task"
)

type taskAndUserQueryService struct {
}

func NewTaskAndUserQueryService() task.TaskAndUserQueryService {
	return &taskAndUserQueryService{}
}

func (qs *taskAndUserQueryService) FetchList(pagingCondition page.PagingCondition) (usecase_task.TaskAndUserDTOPage, error) {
	taskAndUserDTOs := make([]usecase_task.TaskAndUserDTO, 0)

	//TODO tasksとusersをjoinして取得

	page := usecase_task.TaskAndUserDTOPage{
		TaskAndUserDTOs: taskAndUserDTOs,
		Paging: page.Paging{
			TotalCount: 0,
			PageSize:   pagingCondition.PageSize,
			PageNumber: 1,
		},
	}
	return page, nil
}
