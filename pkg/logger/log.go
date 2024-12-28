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
	newKeysAndValues := make([]interface{}, len(l.keysAndValues))
	copy(newKeysAndValues, l.keysAndValues)
	if err != nil {
		newKeysAndValues = append(newKeysAndValues, "err", err)
	}

	l.sugar.Errorw(fmt.Sprintf(msg, a...), newKeysAndValues...)
}

func (l *logger) Warn(msg string, a ...any) {
	l.sugar.Warnw(fmt.Sprintf(msg, a...), l.keysAndValues...)
}

func (l *logger) With(keysAndValues ...interface{}) Logger {
	newKeysAndValues := make([]interface{}, len(l.keysAndValues))
	copy(newKeysAndValues, l.keysAndValues)

	newKeysAndValues = append(newKeysAndValues, keysAndValues...)
	return MakeLogger(newKeysAndValues...)
}
