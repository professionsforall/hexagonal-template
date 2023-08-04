package logger

type AppLogger interface {
	Info(msg string)
	Error(err ...any)
	Panic(err error)
	Fatal(msg string)
}
