package error

type UseCaseError struct {
	message string
}

func NewUseCaseError(message string) UseCaseError {
	return UseCaseError{message: message}
}

func (e UseCaseError) Error() string {
	return e.message
}
