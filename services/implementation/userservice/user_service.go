package userservice

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

type userService struct {
	loggerService interfaces.LoggerService
}

// CreateUser method is responsible for creating a new user.
func (u *userService) CreateUser(name string) shared.User {
	u.loggerService.Info("New user created: %v", name)
	return newUser(name)
}
