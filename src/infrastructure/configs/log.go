package configs

type LogConfigs struct {
	LogLevel string
}

const MissingLogLevelMsg string = `log configurations errors : ( %v ) `

func GetLogConfigs() (*LogConfigs, error) {
	var logConfig, emptyConfig *LogConfigs

	var err error
	logConfig.LogLevel, err = getCompulsoryEnvVar("LOG_LEVEL")
	if err != nil {
		return emptyConfig, err
	}
	return logConfig, nil
}
