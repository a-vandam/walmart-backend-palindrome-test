package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

/*LogContract sets the contract that all loggers passed as dependencies in this project must respect*/
type LogContract interface {
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Error(string, ...interface{})
	Trace(string, ...interface{})
	Fatal(string, ...interface{})
}

/*Log struct implements internally the logging package utilized. In this case, logrus.logger.
TO change it, just change or remove the logger package field, and then reimplement methods accordingly. */
type Log struct {
	loggerPackage *logrus.Logger
}

/*Info wraps the logger package's Info method. Used in standard logging*/
func (l Log) Info(format string, msg ...interface{}) {
	l.loggerPackage.Infof(format, msg...)
}

/*Debug wrraps the logger package Debug. Should be used by criteria*/
func (l Log) Debug(format string, msg ...interface{}) {
	l.loggerPackage.Debugf(format, msg...)
}

/*Error receives a format that has each %v replaced with %+v, to show any stack trace needed. */
func (l Log) Error(format string, msg ...interface{}) {
	formatWithStackTrace := strings.ReplaceAll(format, "%v", "%+v")
	l.loggerPackage.Errorf(formatWithStackTrace, msg...)
}

/*Trace should only be used in sensitive and extensive operation, like logging an extensive SProcedure or Structs as results.*/
func (l Log) Trace(format string, msg ...interface{}) {
	l.loggerPackage.Tracef(format, msg...)
}

/*Fatal should be scarcely used, as it triggers the logger package's Fatal (a combination of Error + OS Exit)*/
func (l Log) Fatal(format string, msg ...interface{}) {
	l.loggerPackage.Fatalf(format, msg...)
}
