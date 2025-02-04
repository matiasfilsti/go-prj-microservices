package controllers

import (
	"authentication/src/api/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Ping(c *gin.Context) {
	s.Core.UserService.GetUser()
	c.String(http.StatusOK, "pong")
}

func (s *Server) SaveUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		RespondHttpError(c, http.StatusBadRequest, err, "Invalid User")
	}

	if err := s.Core.UserService.SaveUser(c, user); err != nil {
		RespondHttpError(c, http.StatusInternalServerError, err, "Error saving user")
	}
	c.JSON(http.StatusOK, user)
}
