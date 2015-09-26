package main

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
)

type libraryApplication struct {
	userService   interfaces.UserService
	bookService   interfaces.BookService
	borrowService interfaces.BorrowService
}

func (l *libraryApplication) run() {
	user := l.userService.CreateUser("Marcin Tojek")
	book := l.bookService.CreateBook("Thinking in Java")
	l.borrowService.Borrow(user, book)
}
