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
	s.Router.GET("/get/:name", s.GetUser)
	s.Router.GET("/authorizeduserjson", s.AuthUserJson)
	s.Router.GET("/authorizeduserbasic", s.AuthUserBasic)
	s.Router.GET("/authorizeduserredis", s.AuthUserLogin)
	s.Router.GET("/loginuserjwt", s.LoginUserJwt)
	s.Router.GET("/authorizeduserjwt", s.AuthUserJwt)

}
