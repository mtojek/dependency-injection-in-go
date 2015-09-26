package borrowingformatter

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
)

func newBorrowingFormatter() interfaces.BorrowingFormatter {
	return new(BorrowingFormatter)
}

func New() interfaces.BorrowingFormatter {
	return newBorrowingFormatter()
}
