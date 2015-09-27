package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

// BorrowingFormatter is an interface which allows formatting borrowings.
type BorrowingFormatter interface {
	Format(user shared.User, book shared.Book) string
}
