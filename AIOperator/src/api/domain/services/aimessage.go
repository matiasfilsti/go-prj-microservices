package services

import (
	"aioperator/src/api/domain/contracts"
	"context"

	"github.com/tmc/langchaingo/schema"
)

type AiMessageService struct {
	repo contracts.AiOperatorRepository
}

func NewAiMessageService(repo contracts.AiOperatorRepository) AiMessageService {
	return AiMessageService{
		repo: repo,
	}
}

func (s *AiMessageService) GenerateResponse(ctx context.Context, docs []schema.Document) (string, error) {
	return s.repo.GenerateResponse(ctx, docs)
}
