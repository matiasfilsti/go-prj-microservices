package controllers

import (
	error "authentication/src/api/domain/errors"
	"authentication/src/api/domain/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AuthUserLogin(c *gin.Context) {
	var name, password string
	var ierr *error.InputError
	var nerr *error.UserNotFoundError

	name, password, ok := c.Request.BasicAuth()
	if !ok {
		RespondHttpError(c, http.StatusInternalServerError, errors.New("basic Auth error"), "Error trying to read user, password")
		return
	}
	user := models.User{
		Name:     name,
		Password: password,
	}
	valid, err := s.Core.UserService.CompareUserPassword(c, user)
	fmt.Println(valid, err)
	if err != nil {
		switch {
		case errors.As(err, &ierr):
			RespondHttpError(c, http.StatusBadRequest, err, "Invalid user information")
			return
		case errors.As(err, &nerr):
			RespondHttpError(c, http.StatusInternalServerError, err, "Invalid name, user not found")
			return
		}
		RespondHttpError(c, http.StatusInternalServerError, err, "Error comparing user, internal-server error")
		return
	}

	if !valid {
		RespondHttpValidUserLogin(c, http.StatusUnauthorized, valid, "", "", "User not allowed")
		return
	}
	sessionToken, csrfToken, err := s.Core.UserService.GenerateTokenUser(c, user.Name)
	if err != nil {
		RespondHttpError(c, http.StatusInternalServerError, err, "error generating token")
		return
	}
	RespondHttpValidUserLogin(c, http.StatusOK, valid, sessionToken, csrfToken, "User allowed")
}
