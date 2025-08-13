package services

import (
	"aioperator/src/api/domain/contracts"
	"context"
)

type AiMessageService struct {
	repo contracts.AiOperatorRepository
}

func NewAiMessageService(repo contracts.AiOperatorRepository) AiMessageService {
	return AiMessageService{
		repo: repo,
	}
}

func (s *AiMessageService) GenerateAIResponse(ctx context.Context, question string) (string, error) {
	return s.repo.GenerateResponse(ctx, question)
}
