package userservice

import (
	"sync"

	"github.com/mtojek/dependency-injection-in-go/services/implementation/loggerservice"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
)

var (
	once     sync.Once
	instance interfaces.UserService
)

func newUserService() interfaces.UserService {
	return &userService{
		loggerService: loggerservice.Instance(),
	}
}

// Instance method returns a singleton instance of UserService.
func Instance() interfaces.UserService {
	once.Do(func() {
		instance = newUserService()
	})
	return instance
}
