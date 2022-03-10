package configs

import "fmt"

type LogConfigs struct {
	LogLevel string
}

var LogConfig *LogConfigs

const MissingLogLevelMsg string = `log configurations errors : ( %v ) `

func GetLogConfigs() (*LogConfigs, error) {
	var reportedErrors []error

	readMustVars.Do(func() {
		LogConfig = new(LogConfigs)
		var err error
		var missingKeysErrors []error
		LogConfig.LogLevel, err = getCompulsoryEnvVar("LOG_LEVEL")
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		reportedErrors = missingKeysErrors
	},
	)
	if len(reportedErrors) != 0 {
		return LogConfig, fmt.Errorf(MissingLogLevelMsg, reportedErrors)
	}

	return LogConfig, nil
}
