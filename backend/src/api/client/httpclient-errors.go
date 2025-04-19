package client

type HttpReqCreationError struct {
	Msg string
}

func (e *HttpReqCreationError) Error() string {
	return e.Msg
}

func NewHttpReqCreationError(msg string) *HttpReqCreationError {
	return &HttpReqCreationError{
		Msg: msg,
	}
}
