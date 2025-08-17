package services

import (
	"aioperator/src/api/domain/contracts"
	"context"

	"github.com/tmc/langchaingo/schema"
)

type RetrievalCacheService struct {
	repo contracts.RetrievalCache
}

func NewRetrievalCacheService(repo contracts.RetrievalCache) RetrievalCacheService {
	return RetrievalCacheService{
		repo: repo,
	}
}

func (s *RetrievalCacheService) Search(ctx context.Context, query string) ([]schema.Document, error) {
	return s.repo.Search(ctx, query)
}
