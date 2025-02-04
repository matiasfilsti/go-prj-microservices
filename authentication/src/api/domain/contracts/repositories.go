package contracts

import (
	"authentication/src/api/domain/models"
	"context"
)

type UserRepository interface {
	Get(ctx context.Context, name string) (*models.User, error)
	Save(ctx context.Context, user models.User) error
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, name string) error
}
