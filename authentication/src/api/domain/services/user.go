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

func (s *UserService) DeleteUser() {

}

func (s *UserService) UpdateUser() {

}

func (s *UserService) ComparePassword() {

}
