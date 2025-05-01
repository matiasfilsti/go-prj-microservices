package controllers

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

type ValidUser struct {
	Valid       bool   `json:"valid"`
	Description string `json:"description"`
}

func RespondHttpValidUser(c *gin.Context, status int, valid bool, desc string) {
	errorResponse := ValidUser{
		Valid:       valid,
		Description: desc,
	}
	c.JSON(status, errorResponse)
}

type ValidUserLogin struct {
	Valid        bool   `json:"valid"`
	Description  string `json:"description"`
	SessionToken string `json:"sessiontoken"`
	CSRFToken    string `json:"csrftoken"`
}

func RespondHttpValidUserLogin(c *gin.Context, status int, valid bool, sessionToken string, csrfToken string, desc string) {
	errorResponse := ValidUserLogin{
		Valid:        valid,
		SessionToken: sessionToken,
		CSRFToken:    csrfToken,
		Description:  desc,
	}
	c.JSON(status, errorResponse)
}

type ValidUserJwt struct {
	Valid       bool   `json:"valid"`
	Description string `json:"description"`
	JwtToken    string `json:"jwttoken"`
}

func RespondHttpValidUserJwt(c *gin.Context, status int, valid bool, jwt string, desc string) {
	errorResponse := ValidUserJwt{
		Valid:       valid,
		JwtToken:    jwt,
		Description: desc,
	}
	c.JSON(status, errorResponse)
}
