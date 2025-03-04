package controllers

import (
	"consumer/src/api/domain"
)

type Server struct {
	RabbitCh *domain.Core
}

func NewServer(core *domain.Core) *Server {
	return &Server{
		RabbitCh: core,
	}
}
