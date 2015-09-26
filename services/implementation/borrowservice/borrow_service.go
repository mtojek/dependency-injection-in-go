package borrowservice

import (
	"github.com/mtojek/dependency-injection-in-go/services/implementation/borrowingformatter"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"
)

type borrowService struct {
	loggerService interfaces.LoggerService
}

// Borrow method is responsible for borrowing book by user.
func (b *borrowService) Borrow(user shared.User, book shared.Book) {
	formatter := borrowingformatter.New()
	formatter.SetUser(user)
	formatter.SetBook(book)
	b.loggerService.Info(formatter.Format())
}
