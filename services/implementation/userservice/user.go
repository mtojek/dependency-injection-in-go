package userservice

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

type user struct {
	name string
}

func newUser(name string) shared.User {
	return &user{name: name}
}

func (u *user) Name() string {
	return u.name
}
