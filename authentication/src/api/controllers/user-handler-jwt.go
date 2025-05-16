package controllers

import (
	error "authentication/src/api/domain/errors"
	"authentication/src/api/domain/models"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) LoginUserJwt(c *gin.Context) {
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
	if err != nil || !valid {
		switch {
		case errors.As(err, &ierr):
			RespondHttpError(c, http.StatusBadRequest, err, "Invalid user information")
			return
		case errors.As(err, &nerr):
			RespondHttpError(c, http.StatusInternalServerError, err, "Invalid name, user not found")
			return
		}
		RespondHttpError(c, http.StatusUnauthorized, err, "user cannot access")
		return
	}

	jwtToken, err := s.Core.UserService.GenerateJwtToken(user.Name)
	if err != nil {
		RespondHttpError(c, http.StatusInternalServerError, err, "error generating jwt")
		return
	}
	RespondHttpValidUserJwt(c, http.StatusOK, valid, jwtToken, "User allowed")
}

func (s *Server) AuthUserJwt(c *gin.Context) {
	var verr *error.JwtParseError
	Completetoken := c.GetHeader("Auth-Header")
	token := strings.TrimPrefix(Completetoken, "Bearer ")
	_, err := s.Core.UserService.ValidateJwtToken(token)
	if err != nil {
		switch {
		case errors.As(err, &verr):
			RespondHttpError(c, http.StatusInternalServerError, err, "error parsing jwt")
			return
		}
		RespondHttpError(c, http.StatusUnauthorized, err, "user not authorized")
		return
	}
	RespondHttpValidUser(c, http.StatusOK, true, "User allowed")
}
