package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type logger struct {
	sugar         *zap.SugaredLogger
	keysAndValues []interface{}
}

func (l *logger) Debug(msg string, a ...any) {
	l.sugar.Debugw(fmt.Sprintf(msg, a...), l.keysAndValues...)
}

func (l *logger) Info(msg string, a ...any) {
	l.sugar.Infow(fmt.Sprintf(msg, a...), l.keysAndValues...)
}

func (l *logger) Error(err error, msg string, a ...any) {
	var keysAndValues []interface{}
	copy(keysAndValues, l.keysAndValues)
	if err != nil {
		keysAndValues = append(keysAndValues, "err", err)
	}
	l.sugar.Errorw(fmt.Sprintf(msg, a...), keysAndValues...)
}

func (l *logger) Warn(msg string, a ...any) {
	l.sugar.Warnw(fmt.Sprintf(msg, a...), l.keysAndValues...)
}

func (l *logger) With(keysAndValues ...interface{}) Logger {
	var newKeysAndValues []interface{}
	copy(newKeysAndValues, l.keysAndValues)

	newKeysAndValues = append(newKeysAndValues, keysAndValues...)
	return &logger{
		sugar:         l.sugar,
		keysAndValues: newKeysAndValues,
	}
}
