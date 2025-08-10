package errors

type RabbitmqMsgError struct {
	Msg string
}

func (e *RabbitmqMsgError) Error() string {
	return e.Msg
}

func NewRabbitmqMsgError(msg string) *RabbitmqMsgError {
	return &RabbitmqMsgError{
		Msg: msg,
	}
}
