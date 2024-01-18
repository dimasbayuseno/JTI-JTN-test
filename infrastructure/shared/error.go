package shared

type NotFoundError struct {
	Message string
}

func (notFoundError NotFoundError) Error() string {
	return notFoundError.Message
}

type UnauthorizedError struct {
	Message string
}

func (unauthorizedError UnauthorizedError) Error() string {
	return unauthorizedError.Message
}

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}
