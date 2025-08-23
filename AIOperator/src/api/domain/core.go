package domain

import (
	"aioperator/src/api/domain/services"
	"aioperator/src/api/repositories"
)

type Core struct {
	Llm   services.AiMessageService
	Cache services.RetrievalCacheService
}

func NewCore() *Core {

	llm, e := repositories.NewAiOperator()
	storeCache := repositories.ConnectRetrievalCache(e)
	repoStoreCache := repositories.NewRetrievalCache(storeCache)
	storeCacheService := services.NewRetrievalCacheService(&repoStoreCache)
	aiMessageService := services.NewAiMessageService(repositories.NewAiOperatorRepository(llm, storeCache))

	return &Core{
		Llm:   aiMessageService,
		Cache: storeCacheService,
	}
}
