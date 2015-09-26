package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

// BorrowingFormatter is an interface which allows formatting borrowings.
type BorrowingFormatter interface {
	SetUser(user shared.User)
	SetBook(book shared.Book)
	Format() string
}
