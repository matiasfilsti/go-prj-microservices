package controllers

import (
	"backend/src/api/client"
	"backend/src/api/domain"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router    *gin.Engine
	RabbitCh  *domain.Core
	Hclient   *client.ClientHttp
	RdsClient *client.ClientRedis
}

func NewServer(router *gin.Engine, core *domain.Core, hclient *client.ClientHttp, rdsClient *client.ClientRedis) *Server {
	return &Server{
		Router:    router,
		RabbitCh:  core,
		Hclient:   hclient,
		RdsClient: rdsClient,
	}
}
