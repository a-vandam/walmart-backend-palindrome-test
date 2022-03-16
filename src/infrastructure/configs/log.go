package configs

type LogConfigs struct {
	LogLevel string
}

const logLevelKey string = "LOG_LEVEL"

func GetLogConfigs() (*LogConfigs, error) {
	var logConfig LogConfigs

	logConfig.LogLevel = getOptionalEnvVar(logLevelKey)

	return &logConfig, nil
}
