package borrowingformatter

import (
	"fmt"

	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

type borrowingFormatter struct {
	user shared.User
	book shared.Book
}

// SetUser method sets user to be used by formatter.
func (b *borrowingFormatter) SetUser(user shared.User) {
	b.user = user
}

// SetBook method sets book to be used by formatter.
func (b *borrowingFormatter) SetBook(book shared.Book) {
	b.book = book
}

// Format method formats the borrowing.
func (b *borrowingFormatter) Format() string {
	return fmt.Sprintf("%v borrowed a book named \"%v\".", b.user.Name(), b.book.Name())
}
