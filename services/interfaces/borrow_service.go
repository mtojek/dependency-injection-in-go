package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

// BorrowService is an interface which allows borrowing books.
type BorrowService interface {
	Borrow(user shared.User, book shared.Book)
}
