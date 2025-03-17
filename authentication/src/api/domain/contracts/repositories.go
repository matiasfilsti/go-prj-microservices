package contracts

import (
	"authentication/src/api/domain/models"
	"context"
)

type UserRepository interface {
	Get(ctx context.Context, name string) (*models.User, error)
	Save(ctx context.Context, user models.User) error
	UpdateToken(ctx context.Context, user models.User, sessiontoken string, crsftoken string) error
	Delete(ctx context.Context, name string) error
	// CompareUserPassword(ctx context.Context, user models.User) error
}
