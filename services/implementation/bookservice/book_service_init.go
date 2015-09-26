package bookservice

import (
	"sync"

	"github.com/mtojek/dependency-injection-in-go/services/implementation/loggerservice"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
)

var (
	once     sync.Once
	instance interfaces.BookService
)

func newBookService() interfaces.BookService {
	return &bookService{
		loggerService: loggerservice.Instance(),
	}
}

// Instance method returns a singleton instance of BookService.
func Instance() interfaces.BookService {
	once.Do(func() {
		instance = newBookService()
	})
	return instance
}
