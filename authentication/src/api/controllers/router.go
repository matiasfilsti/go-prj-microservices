package controllers

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) ConfigureRouter() {

	s.Router.Use(gin.Recovery())
	s.Router.RedirectFixedPath = false
	s.Router.RedirectTrailingSlash = true

	s.Router.GET("/ping", s.Ping)
	s.Router.POST("/save", s.SaveUser)
	s.Router.GET("/allowedusers", s.AllowedUser)
	s.Router.GET("/get/:name", s.GetUser)
}
