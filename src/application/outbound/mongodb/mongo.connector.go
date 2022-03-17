package mongodb

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/infrastructure/configs"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*MongoConnector acts as a wrapper that embeds the mongo client, leaving business logic or more query related logic to it's invokeres (the adapter, for ex)*/
type MongoConnector struct {
	client *mongo.Client
	Log    logger.LogContract
}

/*Connect basically sets up a connection to the DB using DB conn required params. Must be pinged later to check it's feasibility.*/
func (a *MongoConnector) Connect(ctx context.Context, configs *configs.ProductsDBConfigurations) (*MongoConnector, error) {
	a.Log.Info("creating URI for DB connection")
	dbType := "mongodb://"
	userAndPwdString := configs.MongoDatabaseUsername + ":" + configs.MongoDatabasePassword
	hostAndPort := configs.MongoDatabaseHost + ":" + configs.MongoDatabasePort
	dbNamePath := "/" + configs.MongoDatabaseName
	authSource := "?authSource=" + configs.MongoAuthSource
	dbURI := dbType + userAndPwdString + "@" + hostAndPort + dbNamePath + authSource
	a.Log.Debug("db Connection URI: %v", dbURI)

	a.Log.Debug("creating client")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI).SetDirect(true))

	if err != nil {
		a.Log.Error("mongo db conn error: %v", err)
		return &MongoConnector{}, err
	}
	a.Log.Info("connected to productgs mongo db!")
	a.client = client
	a.Log.Info("products db mongo db connector connected succesfully")
	return a, nil
}

/*Ping allows to test if the connection has been established succesfully*/
func (a *MongoConnector) Ping(ctx context.Context) error {
	a.Log.Debug("pinging db...")
	return a.client.Ping(ctx, readpref.Primary())
}

/*Disconnect disconnects the Client from the DB*/
func (a *MongoConnector) Disconnect(ctx context.Context) error {
	a.Log.Debug("disconnecting from DB...")
	return a.client.Disconnect(ctx)
}

/*GetDatabaseNamesUsingFilter allows to retrieve databases found using a filter. Useful for debugging purposes*/
func (a *MongoConnector) GetDatabaseNamesUsingFilter(ctx context.Context, filter bson.D) ([]string, error) {
	list, err := a.client.ListDatabaseNames(ctx, filter, options.ListDatabases().SetNameOnly(true))
	return list, err
}

const idBDKey string = "id"
const brandDBKey string = "brand"
const descriptionDBKey string = "description"
const imageURIDBKey string = "image"
const priceDBKey string = "price"

/*GetFromDatabaseUsingFilter allows the connector to return products from the db and collection name using a filter*/
func (a *MongoConnector) GetFromDatabaseUsingFilter(ctx context.Context, databaseName string, collectionName string, filter bson.D) ([]entities.ProductInfo, error) {
	resultBuffer, err := a.client.Database(databaseName, options.Database()).Collection(collectionName, &options.CollectionOptions{}).Find(ctx, filter, &options.FindOptions{})
	if err != nil {
		a.Log.Error("retrieval of results from database :%v and collection :%v with filter: %+v. error: %v", databaseName, collectionName, filter, err)
		return []entities.ProductInfo{}, err
	}
	var resultFound []bson.D
	resultBuffer.All(ctx, &resultFound)
	defer resultBuffer.Close(context.TODO())
	if err != nil {
		a.Log.Error("mapping of results from db to bson object failed: %v", err)
		return []entities.ProductInfo{}, err
	}
	var prodsFound []entities.ProductInfo

	for i, result := range resultFound {
		a.Log.Debug("mapping result found number: %v", i)
		resultMap := result.Map()
		idAsInt32 := resultMap[idBDKey].(int32)
		idAsInt := int(idAsInt32)
		priceAsInt := resultMap[priceDBKey].(int32)
		priceAsFloat := float32(priceAsInt)

		prodsFound = append(prodsFound, entities.ProductInfo{
			ID:          idAsInt,
			Title:       resultMap[brandDBKey].(string),
			Description: resultMap[descriptionDBKey].(string),
			ImageURL:    resultMap[imageURIDBKey].(string),
			FullPrice:   priceAsFloat,
		})
	}
	return prodsFound, nil
}
