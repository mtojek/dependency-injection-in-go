package loggerservice

import "log"

// LoggerService allows to log various messages.
type LoggerService struct{}

// Info method logs message with info prefix.
func (l *LoggerService) Info(message string, values ...interface{}) {
	log.Printf("[INFO]  "+message+"\n", values...)
}

// Debug method logs messages with debug prefix.
func (l *LoggerService) Debug(message string, values ...interface{}) {
	log.Printf("[DEBUG] "+message+"\n", values...)
}
