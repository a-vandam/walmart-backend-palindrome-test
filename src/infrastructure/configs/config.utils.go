package configs

import (
	"fmt"
	"os"
)

/*MissingEnvVarErrorMsg sets standarized error msg in case the env var is missing*/
const MissingEnvVarErrorMsg string = `missing: "%v" compulsory variable`

func getCompulsoryEnvVar(envVarKey string) (string, error) {
	if envVarKey == "" {
		return "", fmt.Errorf(MissingEnvVarErrorMsg, envVarKey)
	}
	envVarObtained := os.Getenv(envVarKey)
	if envVarObtained == "" {
		return "", fmt.Errorf(MissingEnvVarErrorMsg, envVarKey)
	}
	return envVarObtained, nil
}

func getOptionalEnvVar(envVarKey string) string {
	envVarObtained := os.Getenv(envVarKey)
	return envVarObtained
}
