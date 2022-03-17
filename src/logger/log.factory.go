package logger

import (
	"github.com/sirupsen/logrus"
)

/*LogFactory must be used to create all logs, as this allows for configurations to be cohesive around ALL packages, sth logrus sometimes fail to achieve.*/
type LogFactory struct {
	LogLevel string
}

/*DefaultLogLevel is used  in case it's missing in the Env*/
const DefaultLogLevel string = "info"

/*CreateLog return a log that verifies the logger contract interface. If need of change the library,
just change the Log struct defintion, and then reimplement both the Log methods and this creation of the log */
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
