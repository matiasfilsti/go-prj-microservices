package controllers

import (
	"backend/src/api/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) ConfigureRouter() {

	s.Router.Use(gin.Recovery())
	s.Router.Use(middleware.CORSMiddleware())
	s.Router.RedirectFixedPath = false
	s.Router.RedirectTrailingSlash = true
	autorizedJson := s.Router.Group("/")
	autorizedBasic := s.Router.Group("/")
	autorizedRedis := s.Router.Group("/")
	autorizedjwt := s.Router.Group("/")

	autorizedJson.Use(middleware.AuthorizedUserWithJson(s.Hclient))
	{
		autorizedJson.POST("/producerjson", s.Producer)
	}

	autorizedBasic.Use(middleware.AuthorizedUserWithBasic(s.Hclient))
	{
		autorizedBasic.POST("/producerbasic", s.Producer)
	}
	autorizedRedis.Use(middleware.AuthorizedUserWithRedis(s.RdsClient))
	{
		autorizedRedis.POST("/producerredis", s.Producer)
	}
	autorizedjwt.Use(middleware.AuthorizedUserWithJWT(s.Hclient))
	{
		autorizedjwt.POST("/producerjwt", s.Producer)
	}

	s.Router.GET("/ping", s.Ping)
	s.Router.POST("/loginredis", s.LoginWithRedis)
	s.Router.POST("/authorizedredis", s.AuthorizeRedisLogin)
	s.Router.POST("/logout", s.LogoutHandlerRedis)
	s.Router.GET("/loginjwt", s.LoginWithJwt)

}
