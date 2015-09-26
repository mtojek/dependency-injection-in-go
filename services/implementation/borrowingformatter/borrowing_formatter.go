package borrowingformatter

import (
	"fmt"

	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

type BorrowingFormatter struct {
	user shared.User
	book shared.Book
}

func (b *BorrowingFormatter) SetUser(user shared.User) {
	b.user = user
}

func (b *BorrowingFormatter) SetBook(book shared.Book) {
	b.book = book
}

func (b *BorrowingFormatter) Format() string {
	return fmt.Sprintf("%v borrowed a book named \"%v\".", b.user.Name(), b.book.Name())
}
