package configs

type LogConfigs struct {
	LogLevel string
}

const logLevelKey string = "LOG_LEVEL"

func GetLogConfigs() (*LogConfigs, error) {
	var logConfig LogConfigs
	var err error
	logConfig.LogLevel, err = getCompulsoryEnvVar(logLevelKey)
	if err != nil {
		return &LogConfigs{}, err
	}
	return &logConfig, nil
}
