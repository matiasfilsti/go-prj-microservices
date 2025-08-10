package repositories

import (
	"aioperator/src/api/config"
	"context"
	"log"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
)

func NewRetrievalCache(e *embeddings.EmbedderImpl) *redisvector.Store {
	store, err := redisvector.New(context.Background(),
		redisvector.WithConnectionURL(config.RedisHost),
		redisvector.WithIndexName(config.RedisDB, true),
		redisvector.WithEmbedder(e),
	)
	if err != nil {
		log.Fatal(err)
	}
	return store
}
