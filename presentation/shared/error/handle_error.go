package error

import (
	domain_error "ddd-sample/domain/shared/error"
	usecase_error "ddd-sample/usecase/shared/error"
	"errors"
	"fmt"
)

func HandleError(err error) {
	var domainError domain_error.DomainError
	var useCaseError usecase_error.UseCaseError

	// 適宜処理追加
	if errors.As(err, &domainError) {
		fmt.Println(domainError.Error())
	} else if errors.As(err, &useCaseError) {
		fmt.Println(useCaseError.Error())
	} else {
		fmt.Println(err.Error())
	}
}
