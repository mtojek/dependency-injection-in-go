package loggerservice

import "log"

type LoggerService struct{}

func (l *LoggerService) Info(message string, values ...interface{}) {
	log.Printf("[INFO]  "+message+"\n", values...)
}

func (l *LoggerService) Debug(message string, values ...interface{}) {
	log.Printf("[DEBUG] "+message+"\n", values...)
}
