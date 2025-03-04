package domain

import (
	"consumer/src/api/domain/services"
	"consumer/src/api/repositories"
)

type Core struct {
	MessageService services.MessageService
}

func NewCore() *Core {
	ch := repositories.ConnectRMQ()
	repo := repositories.NewMessageRepository(ch)
	messageService := services.NewMessageService(&repo)
	return &Core{
		MessageService: messageService,
	}
}
