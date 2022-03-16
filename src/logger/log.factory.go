package logger

import (
	"github.com/sirupsen/logrus"
)

type LogFactory struct {
	LogLevel string
}

const DefaultLogLevel string = "info"

func (lf *LogFactory) CreateLog(prefix string) LogContract {
	log := Log{
		logrus.New(),
	}

	log.loggerPackage.SetFormatter(&logrus.JSONFormatter{})
	logLvl, err := logrus.ParseLevel(lf.LogLevel)
	if err != nil {
		lf.LogLevel = DefaultLogLevel
		logLvl, _ = logrus.ParseLevel(DefaultLogLevel)
	}
	log.loggerPackage.SetLevel(logLvl)
	return log
}
