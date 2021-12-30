package user

// ドメインサービス
type UserDuplicationChecker struct {
	userRepository UserRepository
}

func NewUserDuplicationChecker(userRepository UserRepository) UserDuplicationChecker {
	return UserDuplicationChecker{userRepository: userRepository}
}

func (udc *UserDuplicationChecker) Exists(user User) (bool, error) {
	duplicateUser, err := udc.userRepository.FindByName(user.name)
	if err != nil {
		return false, err
	}
	return duplicateUser != nil, nil
}
