package ctrserrors

import "github.com/gin-gonic/gin"

type HttpError struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

func RespondHttpError(c *gin.Context, status int, err error, desc string) {
	errorResponse := HttpError{
		Error:       err.Error(),
		Description: desc,
	}
	c.JSON(status, errorResponse)
}

type ReadBodyResponseError struct {
	Msg string
}

func (e *ReadBodyResponseError) Error() string {
	return e.Msg
}

func NewReadBodyResponseError(msg string) *ReadBodyResponseError {
	return &ReadBodyResponseError{
		Msg: msg,
	}
}

type JsonImportError struct {
	Msg string
}

func (e *JsonImportError) Error() string {
	return e.Msg
}

func NewJsonImportError(msg string) *JsonImportError {
	return &JsonImportError{
		Msg: msg,
	}
}
