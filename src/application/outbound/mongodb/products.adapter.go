package mongodb

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsAdapter struct {
	DBConnector ConnectorContract
	Log         logger.LogContract
}

func (a *ProductsAdapter) GetAllProductsDatabases(ctx context.Context) ([]string, error) {
	a.Log.Info("trying to get all products databases names")

	results, err := a.DBConnector.GetUsingFilter(ctx, bson.D{})
	if err != nil {
		a.Log.Error("failed to obtain all databases names: %v", err)
		return []string{}, err
	}
	a.Log.Debug("result obtained: %+v", results)
	return results, nil

}

//TODO: add implementtation
const ProductsDatabaseName string = "promotions"
const ProductsCollectionName string = "products"

func (a *ProductsAdapter) GetProductById(id int, ctx context.Context) (entities.ProductInfo, error) {
	a.Log.Info("trying to get product with an id :%v, from database: %v", id, ProductsDatabaseName)
	results, err := a.DBConnector.GetFromDatabaseUsingFilter(ctx, ProductsDatabaseName, ProductsCollectionName, bson.D{
		primitive.E{Key: "id", Value: id},
	})
	if err != nil {
		a.Log.Error("failed to obtain all databases names: %v", err)
		return entities.ProductInfo{}, err
	}

	a.Log.Debug("result obtained: %+v", results)
	if len(results) == 0 {
		msg := fmt.Sprintf("no registry for id :%v in database", id)
		a.Log.Error(msg)
		return entities.ProductInfo{}, errors.New(msg)
	}
	return results[0], nil

}
