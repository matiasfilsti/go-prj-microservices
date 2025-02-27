package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Producer(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
