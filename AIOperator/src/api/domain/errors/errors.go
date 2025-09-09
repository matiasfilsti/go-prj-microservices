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

type RedisCacheError struct {
	Msg string
}

func (e *RedisCacheError) Error() string {
	return e.Msg
}

func NewRedisCacheError(msg string) *RedisCacheError {
	return &RedisCacheError{
		Msg: msg,
	}
}
