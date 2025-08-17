package repositories

import (
	"aioperator/src/api/config"
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
)

func NewAiOperator() (llms.Model, *embeddings.EmbedderImpl) {
	llm, err := ollama.New(
		ollama.WithModel("llama3.2:1b"),
		ollama.WithFormat("json"),
		ollama.WithServerURL(config.LlmUrl),
	)
	if err != nil {
		log.Fatal(err)
	}
	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		log.Fatal(err)
	}
	return llm, e
}

type AiOperatorRepository struct {
	llm   llms.Model
	store *redisvector.Store
}

func NewAiOperatorRepository(llm llms.Model, store *redisvector.Store) *AiOperatorRepository {
	return &AiOperatorRepository{
		llm:   llm,
		store: store,
	}
}

func (r *AiOperatorRepository) GenerateResponse(ctx context.Context, docs []schema.Document) (string, error) {
	result, err := chains.Run(ctx, chains.NewRetrievalQAFromLLM(
		r.llm,
		vectorstores.ToRetriever(r.store, 1, vectorstores.WithScoreThreshold(0.5)),
	),
		fmt.Sprintf("genera una respuesta sobre la ciudad:%s y su poblacion: %s millones", docs[0].PageContent, docs[0].Metadata["population"]),
	)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}
