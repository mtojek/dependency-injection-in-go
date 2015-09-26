package interfaces

type LoggerService interface {
	Info(message string, values ...interface{})
	Debug(message string, values ...interface{})
}
