package domain

import (
	"authentication/src/api/domain/services"
	"authentication/src/api/repositories"
	"context"
)

type Core struct {
	UserService services.UserService
}

func NewCore() *Core {
	db := repositories.ConnectDb()
	migrator := repositories.NewMigrationRunner(db)
	migrator.RunMigrations(context.Background())
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)
	return &Core{
		UserService: userService,
	}
}
