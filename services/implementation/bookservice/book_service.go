package bookservice

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

type BookService struct {
	loggerService interfaces.LoggerService
}

func (b *BookService) CreateBook(name string) shared.Book {
	b.loggerService.Info("New book created: %v", name)
	return newBook(name)
}
