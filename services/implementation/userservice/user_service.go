package userservice

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

// UserService allows creating users.
type UserService struct {
	LoggerService interfaces.LoggerService `inject:""`
}

// CreateUser method is responsible for creating a new user.
func (u *UserService) CreateUser(name string) shared.User {
	u.LoggerService.Info("New user created: %v", name)
	return newUser(name)
}
