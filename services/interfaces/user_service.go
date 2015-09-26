package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

// UserService is an interface which allows creating users.
type UserService interface {
	CreateUser(name string) shared.User
}
