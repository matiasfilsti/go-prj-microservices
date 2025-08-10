package controllers

import (
	"aioperator/src/api/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) ConfigureRouter() {

	s.Router.Use(gin.Recovery())
	s.Router.Use(middleware.CORSMiddleware())

	s.Router.GET("/ping", s.Ping)
	s.Router.POST("/ask", s.Ask)

}
