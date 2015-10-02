package main

import (
	"log"

	"github.com/facebookgo/inject"
	"github.com/mtojek/dependency-injection-in-go/services/implementation/bookservice"
	"github.com/mtojek/dependency-injection-in-go/services/implementation/borrowservice"
	"github.com/mtojek/dependency-injection-in-go/services/implementation/loggerservice"
	"github.com/mtojek/dependency-injection-in-go/services/implementation/userservice"
)

func main() {
	application := new(libraryApplication)

	if error := inject.Populate(
		application,
		new(bookservice.BookService),
		new(borrowservice.BorrowService),
		new(loggerservice.LoggerService),
		new(userservice.UserService),
	); nil != error {
		log.Fatalf("Error occured while populating graph: %v", error)
	}

	application.run()
}
