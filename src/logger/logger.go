package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogContract interface {
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Error(string, ...interface{})
	Trace(string, ...interface{})
	Fatal(string, ...interface{})
}

type Log struct {
	loggerPackage *logrus.Logger
}

func (l Log) Info(format string, msg ...interface{}) {
	l.loggerPackage.Infof(format, msg...)
}

func (l Log) Debug(format string, msg ...interface{}) {
	l.loggerPackage.Debugf(format, msg...)
}

//TODO: error now prints stack trace on loggin. Experimental, tbh.
func (l Log) Error(format string, msg ...interface{}) {
	formatWithStackTrace := strings.ReplaceAll(format, "%v", "%+v")
	l.loggerPackage.Errorf(formatWithStackTrace, msg...)
}

func (l Log) Trace(format string, msg ...interface{}) {
	l.loggerPackage.Tracef(format, msg...)
}
func (l Log) Fatal(format string, msg ...interface{}) {
	l.loggerPackage.Fatalf(format, msg...)
}
