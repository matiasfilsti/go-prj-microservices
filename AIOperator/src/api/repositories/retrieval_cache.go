package repositories

import (
	"aioperator/src/api/config"
	"context"
	"log"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
)

func ConnectRetrievalCache(e *embeddings.EmbedderImpl) *redisvector.Store {
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

type RetrievalCache struct {
	store *redisvector.Store
}

func NewRetrievalCache(store *redisvector.Store) RetrievalCache {
	return RetrievalCache{
		store: store,
	}
}

func (r *RetrievalCache) Search(ctx context.Context, query string) ([]string, error) {
	r.store.AddDocuments(ctx, []schema.Document{})
	return nil, nil
}
