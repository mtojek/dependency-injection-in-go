package bookservice

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

// BookService allows to create books.
type BookService struct {
	LoggerService interfaces.LoggerService `inject:""`
}

// CreateBook method is responsible for creating new book.
func (b *BookService) CreateBook(title string) shared.Book {
	b.LoggerService.Info("New book created: %v", title)
	return newBook(title)
}
