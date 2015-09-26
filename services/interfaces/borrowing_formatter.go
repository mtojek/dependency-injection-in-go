package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

type BorrowingFormatter interface {
	SetUser(user shared.User)
	SetBook(book shared.Book)
	Format() string
}
