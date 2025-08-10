package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Ping(c *gin.Context) {
	fmt.Println("Ping endpoint hit")
	c.String(http.StatusOK, "pong")
}
