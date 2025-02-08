package errors

type InputError struct {
	Msg string
}

func (e *InputError) Error() string {
	return e.Msg
}

func NewInputError(msg string) *InputError {
	return &InputError{
		Msg: msg,
	}
}

type ConstraingError struct {
	Msg string
}

func (e *ConstraingError) Error() string {
	return e.Msg
}

func NewConstraingError(msg string) *ConstraingError {
	return &ConstraingError{
		Msg: msg,
	}
}
