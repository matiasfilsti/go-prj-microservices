package contracts

import "context"

type AiOperatorRepository interface {
	GenerateResponse(ctx context.Context, question string) (string, error)
}

type RetrievalCache interface {
	Search(ctx context.Context, query string) ([]string, error)
}
