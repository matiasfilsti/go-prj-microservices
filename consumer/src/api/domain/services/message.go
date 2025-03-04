package services

import (
	"consumer/src/api/domain/contracts"

	"golang.org/x/net/context"
)

type MessageService struct {
	repo contracts.MessageService
}

func NewMessageService(repo contracts.MessageService) MessageService {
	return MessageService{
		repo: repo,
	}
}

func (m *MessageService) SendMessageToQueue(ctx context.Context) {
	m.repo.Send(ctx)
}

func (m *MessageService) ReadMessageFromQueue() {
	m.repo.Read()
}
