package repositories

import (
	"aioperator/src/api/config"
	"context"
	"log"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
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
	llm llms.Model
}

func NewAiOperatorRepository(llm llms.Model) *AiOperatorRepository {
	return &AiOperatorRepository{
		llm: llm,
	}
}

func (r *AiOperatorRepository) GenerateResponse(ctx context.Context, question string) (string, error) {
	return "", nil
}
