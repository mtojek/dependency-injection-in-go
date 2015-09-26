package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

type UserService interface {
	CreateUser(name string) shared.User
}
