package controllers

import (
	error "authentication/src/api/domain/errors"
	"authentication/src/api/domain/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AllowedUserV2(c *gin.Context) {
	var name, password string
	var ierr *error.InputError
	var nerr *error.UserNotFoundError

	name, password, ok := c.Request.BasicAuth()
	if !ok {
		RespondHttpError(c, http.StatusInternalServerError, errors.New("basic Auth error"), "Error trying to read user, password")
		c.Abort()
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
		RespondHttpValidUser(c, http.StatusUnauthorized, valid, "User not allowed")
		return
	}
	RespondHttpValidUser(c, http.StatusOK, valid, "User allowed")
}
