package interfaces

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

type BorrowService interface {
	Borrow(user shared.User, book shared.Book)
}
