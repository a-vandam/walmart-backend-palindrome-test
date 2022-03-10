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

func CreateLog(level string, prefix string) LogContract {
	log := Logger{
		logrus.New(),
	}
	log.loggerPackage.SetFormatter(&logrus.JSONFormatter{})
	switch {
	case strings.EqualFold(level, logrus.InfoLevel.String()):
		log.loggerPackage.SetLevel(logrus.InfoLevel)
	case strings.EqualFold(level, logrus.DebugLevel.String()):
		log.loggerPackage.SetLevel(logrus.DebugLevel)
	case strings.EqualFold(level, logrus.TraceLevel.String()):
		log.loggerPackage.SetLevel(logrus.TraceLevel)
	default:
		log.loggerPackage.SetLevel(logrus.InfoLevel)
	}

	return log

}

type Logger struct {
	loggerPackage *logrus.Logger
}

func (l Logger) Info(format string, msg ...interface{}) {
	l.loggerPackage.Infof(format, msg...)
}

func (l Logger) Debug(format string, msg ...interface{}) {
	l.loggerPackage.Debugf(format, msg...)
}

func (l Logger) Error(format string, msg ...interface{}) {
	l.loggerPackage.Errorf(format, msg...)
}

func (l Logger) Trace(format string, msg ...interface{}) {
	l.loggerPackage.Tracef(format, msg...)
}
func (l Logger) Fatal(format string, msg ...interface{}) {
	l.loggerPackage.Fatalf(format, msg...)
}
