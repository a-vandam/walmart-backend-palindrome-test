package configs

import (
	"fmt"
	"sync"
)

const DBConfigEnvVarsMissingMsg string = "db configurations errors : ( %v )"
const mongoDatabaseNameKey string = "MONGO_DATABASE_NAME"
const mongoDatabasePortKey string = "MONGO_DATABASE_PORT"
const mongoDatabaseUserKey string = "MONGO_DATABASE_USER"
const mongoDatabasePasswordKey string = "MONGO_DATABASE_PASSWORD"
const mongoDatabaseHostKey string = "MONGO_DATABASE_HOST"
const mongoDatabaseAuthSourceKey string = "AUTH_SOURCE"

func GetProductsDatabaseConfigs() (ProductsDBConfigurations, error) {
	var reportedErrors []error
	readMustVars.Do(func() {

		var err error
		var missingKeysErrors []error
		ProductsDBConfig.MongoDatabaseName, err = getCompulsoryEnvVar(mongoDatabaseNameKey)
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		ProductsDBConfig.MongoDatabasePort, err = getCompulsoryEnvVar(mongoDatabasePortKey)
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		ProductsDBConfig.MongoDatabaseUsername, err = getCompulsoryEnvVar(mongoDatabaseUserKey)
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		ProductsDBConfig.MongoDatabasePassword, err = getCompulsoryEnvVar(mongoDatabasePasswordKey)
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		ProductsDBConfig.MongoDatabaseHost, err = getCompulsoryEnvVar(mongoDatabaseHostKey)
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		ProductsDBConfig.MongoAuthSource, err = getCompulsoryEnvVar(mongoDatabaseAuthSourceKey)
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		reportedErrors = missingKeysErrors
	},
	)
	if len(reportedErrors) != 0 {
		return ProductsDBConfig, fmt.Errorf(DBConfigEnvVarsMissingMsg, reportedErrors)
	}
	return ProductsDBConfig, nil
}

type ProductsDBConfigurations struct {
	MongoDatabaseName     string
	MongoDatabasePort     string
	MongoDatabaseHost     string
	MongoDatabaseUsername string
	MongoDatabasePassword string
	MongoAuthSource       string
}

var readMustVars sync.Once

var ProductsDBConfig ProductsDBConfigurations
