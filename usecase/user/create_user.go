package user

import domain "ddd-sample/domain/user"

type CreateUserUseCase interface {
	Execute(input CreateUserUseCaseInput) (CreateUserUseCaseOutput, error)
}

type CreateUserUseCaseInput struct { // TODO 名称検討
	Name string
}

type CreateUserUseCaseOutput struct { // TODO 名称検討
}

type createUserUseCase struct {
	userRepository domain.UserRepository
	// TODO domain service
}

func (uc *createUserUseCase) Execute(input CreateUserUseCaseInput) (CreateUserUseCaseOutput, error) {
	user := domain.NewUser(input.Name)
	err := uc.userRepository.Insert(user)
	if err != nil {
		return CreateUserUseCaseOutput{}, err
	}
	return CreateUserUseCaseOutput{}, nil
}

func NewCreateUserUseCase(userRepository domain.UserRepository) CreateUserUseCase {
	return &createUserUseCase{
		userRepository: userRepository,
	}
}
