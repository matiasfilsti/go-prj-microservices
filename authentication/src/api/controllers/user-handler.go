package controllers

import (
	error "authentication/src/api/domain/errors"
	"authentication/src/api/domain/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Ping(c *gin.Context) {
	s.Core.UserService.GetUser()
	c.String(http.StatusOK, "pong")
}

func (s *Server) SaveUser(c *gin.Context) {
	var user models.User
	var ierr *error.InputError
	if err := c.ShouldBindJSON(&user); err != nil {
		RespondHttpError(c, http.StatusBadRequest, err, "Invalid Body Data")
	}
	fmt.Println(user)
	if err := s.Core.UserService.SaveUser(c, user); err != nil {
		switch {
		case errors.As(err, &ierr):
			RespondHttpError(c, http.StatusBadRequest, err, "Invalid user information, user or password doesnt meet the requirements")
		}
		RespondHttpError(c, http.StatusInternalServerError, err, "Error saving user, internal-server error")
	}
	c.JSON(http.StatusOK, user)
}
