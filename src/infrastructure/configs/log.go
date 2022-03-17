package configs

/*LogConfigs is a struct made for passing around Log or logging configurations*/
type LogConfigs struct {
	LogLevel string
}

const logLevelKey string = "LOG_LEVEL"

/*GetLogConfigs Returns log configurations. Making them compulsory or optional should be added here. */
func GetLogConfigs() (*LogConfigs, error) {
	var logConfig LogConfigs

	logConfig.LogLevel = getOptionalEnvVar(logLevelKey)

	return &logConfig, nil
}
