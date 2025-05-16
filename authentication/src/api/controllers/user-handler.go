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
	c.String(http.StatusOK, "pong")
}

func (s *Server) SaveUser(c *gin.Context) {
	var user models.User
	var ierr *error.InputError
	var cerr *error.ConstraingError
	var eerr *error.PasswordHashError
	if err := c.ShouldBindJSON(&user); err != nil {
		RespondHttpError(c, http.StatusBadRequest, err, "Invalid Body Data")
	}
	fmt.Println(user)
	if err := s.Core.UserService.SaveUser(c, user); err != nil {
		switch {
		case errors.As(err, &ierr):
			RespondHttpError(c, http.StatusBadRequest, err, "Invalid user information, user or password doesnt meet the requirements")
			return
		case errors.As(err, &cerr):
			RespondHttpError(c, http.StatusBadRequest, err, "Constrain error, user already present")
			return
		case errors.As(err, &eerr):
			RespondHttpError(c, http.StatusInternalServerError, err, "Internal server error, password encryption problem")
			return
		}

		RespondHttpError(c, http.StatusInternalServerError, err, "Error saving user, internal-server error")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) GetUser(c *gin.Context) {
	name := c.Param("name")
	var nerr *error.UserNotFoundError
	var userdb *models.User
	userdb, err := s.Core.UserService.GetUser(c, name)
	if err != nil {
		switch {
		case errors.As(err, &nerr):
			RespondHttpError(c, http.StatusNotFound, err, "Invalid name, user not found")
			return
		}

		RespondHttpError(c, http.StatusInternalServerError, err, "Error comparing user, internal-server error")
		return
	}
	c.JSON(http.StatusOK, userdb)
}
