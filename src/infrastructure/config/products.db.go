package configs

import (
	"fmt"
	"sync"
)

type ProductsDBConfigurations struct {
	DATABASE_NAME     string
	DATABASE_PORT     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	AUTH_SOURCE       string
}

var readMustVars sync.Once
var ProductsDBConfig ProductsDBConfigurations

func init() {

}

const DBConfigEnvVarsMissingMsg string = "the following are the names of missing env vars: ( %v )"

func GetProductsDatabaseConfigs() (ProductsDBConfigurations, error) {
	var reportedErrors []error

	readMustVars.Do(func() {
		ProductsDBConfig := new(ProductsDBConfigurations)
		var err error
		var missingKeysErrors []error
		ProductsDBConfig.DATABASE_NAME, err = getCompulsoryEnvVar("DATABASE_NAME")
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)

		}
		ProductsDBConfig.DATABASE_PORT, err = getCompulsoryEnvVar("DATABASE_PORT")
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		ProductsDBConfig.DATABASE_NAME, err = getCompulsoryEnvVar("DATABASE_NAME")
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		ProductsDBConfig.DATABASE_PASSWORD, err = getCompulsoryEnvVar("DATABASE_PASSWORD")
		if err != nil {
			missingKeysErrors = append(missingKeysErrors, err)
		}
		ProductsDBConfig.AUTH_SOURCE, err = getCompulsoryEnvVar("AUTH_SOURCE")
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
