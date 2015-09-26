package borrowingformatter

import (
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
)

func newBorrowingFormatter() interfaces.BorrowingFormatter {
	return new(borrowingFormatter)
}

// New method returns a new instance of borrowingFormatter
func New() interfaces.BorrowingFormatter {
	return newBorrowingFormatter()
}
