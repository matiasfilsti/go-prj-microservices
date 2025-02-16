package controllers

import "github.com/gin-gonic/gin"

func RespondHttpError(c *gin.Context, status int, err error, desc string) {
	errorResponse := HttpError{
		Error:       err.Error(),
		Description: desc,
	}
	c.JSON(status, errorResponse)
}

type HttpError struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

func RespondHttpValidUser(c *gin.Context, status int, valid bool, desc string) {
	errorResponse := ValidUser{
		Valid:       valid,
		Description: desc,
	}
	c.JSON(status, errorResponse)
}

type ValidUser struct {
	Valid       bool   `json:"valid"`
	Description string `json:"description"`
}
