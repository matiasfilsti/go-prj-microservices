package domain

import (
	"aioperator/src/api/repositories"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
)

type Core struct {
	llm   llms.Model
	cache *redisvector.Store
}

func NewCore() *Core {

	llm, e := repositories.NewAiOperator()
	store := repositories.NewRetrievalCache(e)

	return &Core{
		llm:   llm,
		cache: store,
	}
}
