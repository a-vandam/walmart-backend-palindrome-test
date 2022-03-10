package configs

const mongoDatabaseNameKey string = "MONGO_DATABASE_NAME"
const mongoDatabasePortKey string = "MONGO_DATABASE_PORT"
const mongoDatabaseUserKey string = "MONGO_DATABASE_USER"
const mongoDatabasePasswordKey string = "MONGO_DATABASE_PASSWORD"
const mongoDatabaseHostKey string = "MONGO_DATABASE_HOST"
const mongoDatabaseAuthSourceKey string = "AUTH_SOURCE"

func GetProductsDatabaseConfigs() (*ProductsDBConfigurations, error) {
	var err error
	var emptyConfig ProductsDBConfigurations
	ProductsDBConfig.MongoDatabaseName, err = getCompulsoryEnvVar(mongoDatabaseNameKey)
	if err != nil {
		return &emptyConfig, err
	}
	ProductsDBConfig.MongoDatabasePort, err = getCompulsoryEnvVar(mongoDatabasePortKey)
	if err != nil {
		return &emptyConfig, err
	}
	ProductsDBConfig.MongoDatabaseUsername, err = getCompulsoryEnvVar(mongoDatabaseUserKey)
	if err != nil {
		return &emptyConfig, err
	}
	ProductsDBConfig.MongoDatabasePassword, err = getCompulsoryEnvVar(mongoDatabasePasswordKey)
	if err != nil {
		return &emptyConfig, err
	}
	ProductsDBConfig.MongoDatabaseHost, err = getCompulsoryEnvVar(mongoDatabaseHostKey)
	if err != nil {
		return &emptyConfig, err

	}
	ProductsDBConfig.MongoAuthSource, err = getCompulsoryEnvVar(mongoDatabaseAuthSourceKey)
	if err != nil {
		return &emptyConfig, err
	}

	return &ProductsDBConfig, nil
}

type ProductsDBConfigurations struct {
	MongoDatabaseName     string
	MongoDatabasePort     string
	MongoDatabaseHost     string
	MongoDatabaseUsername string
	MongoDatabasePassword string
	MongoAuthSource       string
}

var ProductsDBConfig ProductsDBConfigurations
