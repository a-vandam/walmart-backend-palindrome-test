package logger

import (
	"github.com/sirupsen/logrus"
)

type LogContract interface {
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Error(string, ...interface{})
	Trace(string, ...interface{})
	Fatal(string, ...interface{})
}

type log struct {
	loggerPackage *logrus.Logger
}

func (l log) Info(format string, msg ...interface{}) {
	l.loggerPackage.Infof(format, msg...)
}

func (l log) Debug(format string, msg ...interface{}) {
	l.loggerPackage.Debugf(format, msg...)
}

func (l log) Error(format string, msg ...interface{}) {
	l.loggerPackage.Errorf(format, msg...)
}

func (l log) Trace(format string, msg ...interface{}) {
	l.loggerPackage.Tracef(format, msg...)
}
func (l log) Fatal(format string, msg ...interface{}) {
	l.loggerPackage.Fatalf(format, msg...)
}
