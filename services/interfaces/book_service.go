package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

type BookService interface {
	CreateBook(name string) shared.Book
}
