package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

// BookService is an interface which allows creating books.
type BookService interface {
	CreateBook(name string) shared.Book
}
