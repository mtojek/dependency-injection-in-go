package loggerservice

import "log"

type loggerService struct{}

func (l *loggerService) Info(message string, values ...interface{}) {
	log.Printf("[INFO]  "+message+"\n", values...)
}

func (l *loggerService) Debug(message string, values ...interface{}) {
	log.Printf("[DEBUG] "+message+"\n", values...)
}
