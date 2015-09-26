package userservice

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

type UserService struct {
	loggerService interfaces.LoggerService
}

func (u *UserService) CreateUser(name string) shared.User {
	u.loggerService.Info("New user created: %v", name)
	return newUser(name)
}
