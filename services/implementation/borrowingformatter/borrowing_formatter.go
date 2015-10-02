package borrowingformatter

import (
	"fmt"

	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

// BorrowingFormatter is responsible for formatting borrowings.
type BorrowingFormatter struct{}

// Format method formats the borrowing.
func (b *BorrowingFormatter) Format(user shared.User, book shared.Book) string {
	return fmt.Sprintf("%v borrowed a book named \"%v\".", user.Name(), book.Title())
}
