package shared

type InvalidInputError struct {
	Msg string
}

type Unauthorized struct {
	Msg string
}

func (e *InvalidInputError) Error() string {
	return e.Msg
}

func NewInvalidInputError(msg string) error {
	return &InvalidInputError{Msg: msg}
}

func (e *Unauthorized) Error() string {
	return e.Msg
}

func NewUnauthorized(msg string) error {
	return &Unauthorized{Msg: "Not authorized: " + msg}
}
