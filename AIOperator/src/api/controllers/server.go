package controllers

import (
	"aioperator/src/api/domain"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Core   *domain.Core
}

func NewServer(router *gin.Engine, core *domain.Core) *Server {
	return &Server{
		Router: router,
		Core:   core,
	}
}
