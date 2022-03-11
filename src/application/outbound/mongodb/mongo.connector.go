package mongodb

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/infrastructure/configs"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnector struct {
	client *mongo.Client
	log    logger.LogContract
}

/*Connect basically connects to DB using DB conn required params*/
func (a *MongoConnector) ConnectAndPing(ctx context.Context, configs *configs.ProductsDBConfigurations) error {
	a.log.Info("creating URI for DB connection")
	dbType := "mongoDB://"
	userAndPwdString := configs.MongoDatabaseUsername + ":" + configs.MongoDatabasePassword
	hostAndPort := configs.MongoDatabaseHost + ":" + configs.MongoDatabasePort
	dbNamePath := "/" + configs.MongoDatabaseName
	authSource := "?authSource=password"
	dbURI := dbType + userAndPwdString + "@" + hostAndPort + dbNamePath + authSource
	a.log.Debug("URI: %v", dbURI)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		a.log.Error("mongo db conn error: %v", err)
		return err
	}
	a.log.Info("connected to productgs mongo db!")
	a.log.Debug("pinging DB")
	err = client.Ping(ctx, nil)
	if err != nil {
		a.log.Error("db ping error: %v", err)
		return err
	}
	a.client = client
	a.log.Info("products db adapter connected succesfully")
	return nil
}

func (a *MongoConnector) Disconnect(ctx context.Context) error {
	a.log.Debug("disconnecting from DB...")
	return a.client.Disconnect(ctx)
}
