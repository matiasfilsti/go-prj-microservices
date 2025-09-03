package contracts

import (
	"context"

	"github.com/tmc/langchaingo/schema"
)

type AiOperatorRepository interface {
	GenerateResponse(ctx context.Context, docs []schema.Document, question string) (string, error)
}

type RetrievalCache interface {
	Search(ctx context.Context, query string) ([]schema.Document, error)
}
