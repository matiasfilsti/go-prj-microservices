package services

import (
	"authentication/src/api/domain/contracts"
	"authentication/src/api/domain/models"
	"context"
)

type UserService struct {
	repo contracts.UserRepository
}

func NewUserService(repo contracts.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser() {

}

func (s *UserService) SaveUser(ctx context.Context, user models.User) error {
	err := userValidate(user)
	if err != nil {
		return err
	}
	return s.repo.Save(ctx, user)
}

func (s *UserService) DeleteUser() {

}

func (s *UserService) UpdateUser() {

}

func (s *UserService) ComparePassword() {

}
