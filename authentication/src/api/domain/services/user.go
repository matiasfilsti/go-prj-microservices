package services

import (
	"authentication/src/api/domain/contracts"
	"authentication/src/api/domain/models"
	"context"
	"fmt"
)

type UserService struct {
	repo contracts.UserRepository
}

func NewUserService(repo contracts.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser(ctx context.Context, username string) (*models.User, error) {
	return s.repo.Get(ctx, username)
}

func (s *UserService) SaveUser(ctx context.Context, user models.User) error {
	err := userValidate(user)
	if err != nil {
		return err
	}
	encriptedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	userEncrypted := models.User{
		Name:     user.Name,
		Password: encriptedPassword,
	}

	return s.repo.Save(ctx, userEncrypted)
}

func (s *UserService) GenerateTokenUser(ctx context.Context, username string) (string, string, error) {
	sessionToken := GenerateToken()
	csrfToken := GenerateToken()
	dbUser, err := s.repo.Get(ctx, username)
	if err != nil {
		return "", "", err
	}
	err = s.repo.UpdateToken(ctx, *dbUser, sessionToken, csrfToken)
	if err != nil {
		return "", "", err
	}
	return sessionToken, csrfToken, nil
}

func (s *UserService) CompareUserPassword(ctx context.Context, user models.User) (bool, error) {
	err := userValidate(user)
	if err != nil {
		return false, err
	}
	dbUserEncripted, err := s.repo.Get(ctx, user.Name)
	if err != nil {
		return false, err
	}
	fmt.Println(dbUserEncripted, user)
	return checkPasswordHash(user.Password, dbUserEncripted.Password), nil

}
