package borrowservice

import (
	"sync"

	"github.com/mtojek/dependency-injection-in-go/services/implementation/loggerservice"
	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
)

var (
	once     sync.Once
	instance interfaces.BorrowService
)

func newBorrowService() interfaces.BorrowService {
	return &BorrowService{
		loggerService: loggerservice.Instance(),
	}
}

func Instance() interfaces.BorrowService {
	once.Do(func() {
		instance = newBorrowService()
	})
	return instance
}
