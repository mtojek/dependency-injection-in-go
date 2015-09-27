package borrowingformatter

import (
	"fmt"

	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

type borrowingFormatter struct{}

// Format method formats the borrowing.
func (b *borrowingFormatter) Format(user shared.User, book shared.Book) string {
	return fmt.Sprintf("%v borrowed a book named \"%v\".", user.Name(), book.Name())
}
