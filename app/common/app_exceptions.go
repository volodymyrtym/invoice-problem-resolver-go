package common

type InvalidInputError struct {
	Msg string
}

func (e *InvalidInputError) Error() string {
	return e.Msg
}

func NewInvalidInputError(msg string) error {
	return &InvalidInputError{Msg: msg}
}
