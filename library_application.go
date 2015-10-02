package main

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
)

type libraryApplication struct {
	UserService   interfaces.UserService   `inject:""`
	BookService   interfaces.BookService   `inject:""`
	BorrowService interfaces.BorrowService `inject:""`
}

func (l *libraryApplication) run() {
	user := l.UserService.CreateUser("Marcin Tojek")
	book := l.BookService.CreateBook("Thinking in Java")
	l.BorrowService.Borrow(user, book)
}
