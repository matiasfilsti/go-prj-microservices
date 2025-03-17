package services

import (
	"backend/src/api/domain/contracts"

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

func (m *MessageService) SendMessageToQueue(ctx context.Context) error {
	return m.repo.Send(ctx)
}
