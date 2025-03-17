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
	autorizedV2 := s.Router.Group("/")
	autorizedV3 := s.Router.Group("/")

	autorized.Use(middleware.AuthorizedUser(s.Hclient))
	{
		autorized.POST("/producer", s.Producer)
	}

	autorizedV2.Use(middleware.AuthorizedUserV2(s.Hclient))
	{
		autorizedV2.POST("/producerV2", s.Producer)
	}
	autorizedV3.Use(middleware.AuthorizedUserV3())
	{
		autorizedV2.POST("/producerV3", s.Producer)
	}

	s.Router.GET("/ping", s.Ping)

	s.Router.GET("/login", s.Login)
	// s.Router.GET("/producer-secured", s.ProducerSecured)

}
