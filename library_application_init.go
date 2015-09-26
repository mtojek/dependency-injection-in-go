package main

import (
	"github.com/mtojek/dependency-injection-in-go/services/implementation/bookservice"
	"github.com/mtojek/dependency-injection-in-go/services/implementation/borrowservice"
	"github.com/mtojek/dependency-injection-in-go/services/implementation/userservice"
)

func newLibraryApplication() *libraryApplication {
	return &libraryApplication{
		userService:   userservice.Instance(),
		bookService:   bookservice.Instance(),
		borrowService: borrowservice.Instance(),
	}
}
