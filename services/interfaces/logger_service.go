package interfaces

// LoggerService is an interface which allows printing important messages.
type LoggerService interface {
	Info(message string, values ...interface{})
	Debug(message string, values ...interface{})
}
