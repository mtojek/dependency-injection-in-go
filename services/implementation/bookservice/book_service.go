package bookservice

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

type bookService struct {
	loggerService interfaces.LoggerService
}

func (b *bookService) CreateBook(title string) shared.Book {
	b.loggerService.Info("New book created: %v", title)
	return newBook(title)
}
