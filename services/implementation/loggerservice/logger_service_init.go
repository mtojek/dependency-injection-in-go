package loggerservice

import (
	"sync"

	"github.com/mtojek/dependency-injection-in-go/services/interfaces"
)

var (
	once     sync.Once
	instance interfaces.LoggerService
)

func newLoggerService() interfaces.LoggerService {
	return new(loggerService)
}

// Instance method returns a singleton instance of LoggerService.
func Instance() interfaces.LoggerService {
	once.Do(func() {
		instance = newLoggerService()
	})
	return instance
}
