package bookservice

type mockedLogger struct {
	invoked bool
}

func (m *mockedLogger) Info(message string, values ...interface{}) {
	m.invoked = true
}

func (m *mockedLogger) Debug(message string, values ...interface{}) {
}
