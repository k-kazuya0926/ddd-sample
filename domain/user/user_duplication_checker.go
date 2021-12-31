package user

import "context"

// ドメインサービス
type UserDuplicationChecker struct {
	userRepository UserRepository
}

func NewUserDuplicationChecker(userRepository UserRepository) UserDuplicationChecker {
	return UserDuplicationChecker{userRepository: userRepository}
}

func (udc *UserDuplicationChecker) Exists(ctx context.Context, user User) (bool, error) {
	duplicateUser, err := udc.userRepository.FindByName(ctx, user.name)
	if err != nil {
		return false, err
	}
	return duplicateUser != nil, nil
}
