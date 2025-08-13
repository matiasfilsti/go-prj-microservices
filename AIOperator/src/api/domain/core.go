package domain

import (
	"aioperator/src/api/domain/services"
	"aioperator/src/api/repositories"

	"github.com/tmc/langchaingo/llms"
)

type Core struct {
	llm   llms.Model
	cache services.RetrievalCacheService
}

func NewCore() *Core {

	llm, e := repositories.NewAiOperator()
	storeCache := repositories.ConnectRetrievalCache(e)
	repoStoreCache := repositories.NewRetrievalCache(storeCache)
	storeCacheService := services.NewRetrievalCacheService(&repoStoreCache)

	return &Core{
		llm:   llm,
		cache: storeCacheService,
	}
}
