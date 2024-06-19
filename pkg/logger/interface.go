package logger

import "go.uber.org/zap"

type Logger interface {
	Debug(msg string, a ...any)
	Info(msg string, a ...any)
	Error(err error, msg string, a ...any)
	Warn(msg string, a ...any)
	With(keysAndValues ...interface{}) Logger
}

func MakeLogger(keysAndValues ...interface{}) Logger {
	l, _ := zap.NewProduction(zap.AddCallerSkip(1))

	return &logger{
		sugar:         l.Sugar(),
		keysAndValues: keysAndValues,
	}

}
