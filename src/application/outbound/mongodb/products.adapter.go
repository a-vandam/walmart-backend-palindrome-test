package adapters

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/infrastructure/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductsAdapter struct {
}

/*Connect basically connects to DB using DB conn required params*/
func (adapter *ProductsAdapter) Connect(ctx context.Context, configs *configs.ProductsDBConfigurations) error {
	dbType := "mongoDB://"
	userAndPwdString := configs.MongoDatabaseUsername + ":" + configs.MongoDatabasePassword
	hostAndPort := configs.MongoDatabaseHost + ":" + configs.MongoDatabasePort
	dbNamePath := "/" + configs.MongoDatabaseName
	authSource := "?authSource=password"
	dbURI := dbType + userAndPwdString + "@" + hostAndPort + dbNamePath + authSource
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {

	}
	return nil
}

var pDBLog logger.LogI

func init() {

}
