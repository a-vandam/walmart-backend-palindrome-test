package configs

import (
	"fmt"
	"os"
)

const MissingEnvVarErrorMsg string = `missing: "%v" compulsory variable`
const NoKeyRetrieved string = "missing environment variable name or key"

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
