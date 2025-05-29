package logger

type Field struct {
	Key   string
	Value interface{}
}

type Logger interface {
	Error(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
}
