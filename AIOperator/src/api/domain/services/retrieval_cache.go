package services

import (
	"aioperator/src/api/domain/contracts"
	"context"
)

type RetrievalCacheService struct {
	repo contracts.RetrievalCache
}

func NewRetrievalCacheService(repo contracts.RetrievalCache) RetrievalCacheService {
	return RetrievalCacheService{
		repo: repo,
	}
}

func (s *RetrievalCacheService) Search(ctx context.Context, query string) ([]string, error) {
	return s.repo.Search(ctx, query)
}
