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
	return &BookService{
		loggerService: loggerservice.Instance(),
	}
}

func Instance() interfaces.BookService {
	once.Do(func() {
		instance = newBookService()
	})
	return instance
}
