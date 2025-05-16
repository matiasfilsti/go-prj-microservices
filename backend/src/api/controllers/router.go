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
	autorizedjwt := s.Router.Group("/")

	autorized.Use(middleware.AuthorizedUserWithJson(s.Hclient))
	{
		autorized.POST("/producerjson", s.Producer)
	}

	autorizedV2.Use(middleware.AuthorizedUserWithBasic(s.Hclient))
	{
		autorizedV2.POST("/producerbasic", s.Producer)
	}
	autorizedV3.Use(middleware.AuthorizedUserWithRedis(s.RdsClient))
	{
		autorizedV3.POST("/producerredis", s.Producer)
	}
	autorizedjwt.Use(middleware.AuthorizedUserWithJWT(s.Hclient))
	{
		autorizedjwt.POST("/producerjwt", s.Producer)
	}

	s.Router.GET("/ping", s.Ping)
	s.Router.GET("/loginredis", s.LoginWithRedis)
	s.Router.GET("/loginjwt", s.LoginWithJwt)

}
