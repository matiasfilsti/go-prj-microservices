package controllers

import (
	error "authentication/src/api/domain/errors"
	"authentication/src/api/domain/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AuthUserJson(c *gin.Context) {
	var user models.User
	var ierr *error.InputError
	var nerr *error.UserNotFoundError

	if err := c.ShouldBindJSON(&user); err != nil {
		RespondHttpError(c, http.StatusBadRequest, err, "Invalid Body Data")
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
