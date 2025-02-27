package controllers

import (
	"backend/src/api/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) ConfigureRouter() {

	s.Router.Use(gin.Recovery())
	s.Router.RedirectFixedPath = false
	s.Router.RedirectTrailingSlash = true
	autorized := s.Router.Group("/")

	autorized.Use(middleware.AuthorizedUser())
	{
		autorized.POST("/producer", s.Producer)
	}
	s.Router.GET("/ping", s.Ping)

}
