package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogFactory struct {
	LogLevel string
}

func (lf *LogFactory) CreateLog(prefix string) LogContract {
	log := log{
		logrus.New(),
	}
	log.loggerPackage.SetFormatter(&logrus.JSONFormatter{})
	switch {
	case strings.EqualFold(lf.LogLevel, logrus.InfoLevel.String()):
		log.loggerPackage.SetLevel(logrus.InfoLevel)
	case strings.EqualFold(lf.LogLevel, logrus.DebugLevel.String()):
		log.loggerPackage.SetLevel(logrus.DebugLevel)
	case strings.EqualFold(lf.LogLevel, logrus.TraceLevel.String()):
		log.loggerPackage.SetLevel(logrus.TraceLevel)
	default:
		log.loggerPackage.SetLevel(logrus.InfoLevel)
	}

	return log

}
