package borrowservice

import (
	"github.com/mtojek/dependency-injection-in-go/services/implementation/borrowingformatter"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

// BorrowService allows to borrow books.
type BorrowService struct {
	LoggerService interfaces.LoggerService `inject:""`
}

// Borrow method is responsible for borrowing book by user.
func (b *BorrowService) Borrow(user shared.User, book shared.Book) {
	formatter := new(borrowingformatter.BorrowingFormatter)
	formatted := formatter.Format(user, book)

	b.LoggerService.Info(formatted)
}
