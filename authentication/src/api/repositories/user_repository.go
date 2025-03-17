package repositories

import (
	errorcustom "authentication/src/api/domain/errors"
	"authentication/src/api/domain/models"
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Get(ctx context.Context, name string) (*models.User, error) {
	s := &models.User{}
	err := r.db.NewSelect().Model(s).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorcustom.NewUserNotFoundError("user not found")
		}
		return nil, err
	}
	return s, nil
}

func (r *UserRepository) Save(ctx context.Context, user models.User) error {
	_, err := r.db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		if err, ok := err.(pgdriver.Error); ok && err.IntegrityViolation() {
			errorcustom.NewConstraingError("problem inserting data, error contrain")
		}
		return err
	}
	return nil
}

func (r *UserRepository) UpdateToken(ctx context.Context, user models.User, sessiontoken string, crsftoken string) error {
	_, err := r.db.NewUpdate().Model(&user).Set("session_token = ?", sessiontoken).Set("csrf_token = ?", crsftoken).Where("name = ?", user.Name).Exec(ctx)
	if err != nil {
		return errorcustom.NewUserTokenUpdateError("error updating token in database")
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, name string) error {

	return nil
}
