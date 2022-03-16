package mongodb

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/infrastructure/configs"
	"go.mongodb.org/mongo-driver/bson"
)

type ConnectorContract interface {
	Connect(ctx context.Context, configs *configs.ProductsDBConfigurations) (*MongoConnector, error)
	Ping(ctx context.Context) error
	Disconnect(ctx context.Context) error
	GetUsingFilter(ctx context.Context, filter bson.D) ([]string, error)
	GetFromDatabaseUsingFilter(ctx context.Context, databaseName string, collectionName string, filter bson.D) ([]entities.ProductInfo, error)
}
