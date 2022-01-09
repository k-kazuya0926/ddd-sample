package error

type DomainError struct {
	message string
}

func NewDomainError(message string) DomainError {
	return DomainError{message: message}
}

func (e DomainError) Error() string {
	return e.message
}
