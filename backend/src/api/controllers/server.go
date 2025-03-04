package controllers

import (
	"backend/src/api/domain"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router   *gin.Engine
	RabbitCh *domain.Core
}

func NewServer(router *gin.Engine, core *domain.Core) *Server {
	return &Server{
		Router:   router,
		RabbitCh: core,
	}
}
