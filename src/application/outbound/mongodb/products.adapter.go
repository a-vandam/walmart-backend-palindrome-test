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

/*ProductsAdapter wraps the Mongo Client to allow decoupling between the client's specifications or configs,
and the use of filters, criterias, optional lookup parameters, etc, that may be needed.*/
type ProductsAdapter struct {
	DBConnector MongoClientContract
	Log         logger.LogContract
}

/*GetAllProductsDatabases queries the DB and returns the databases found. Useful for debugging*/
func (a *ProductsAdapter) GetAllProductsDatabases(ctx context.Context) ([]string, error) {
	a.Log.Info("trying to get all products databases names")

	results, err := a.DBConnector.GetDatabaseNamesUsingFilter(ctx, bson.D{})
	if err != nil {
		a.Log.Error("failed to obtain all databases names: %v", err)
		return []string{}, err
	}
	a.Log.Debug("result obtained: %+v", results)
	return results, nil

}

/*ProductsDatabaseName is the DB name where products collections are stored*/
const ProductsDatabaseName string = "promotions"

/*ProductsCollectionName is the DB Collection where products are stored*/
const ProductsCollectionName string = "products"

/*GetProductByID allows for retrieval of a product that matched the ID sent. If no prod are found, an error is returned.
Context is passed directly to DB CLient.*/
func (a *ProductsAdapter) GetProductByID(ctx context.Context, id int) (entities.ProductInfo, error) {
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

/*GetProductsByText queries the DB with an OR filter in brand or description fields. Those must mach the text param sent*/
func (a *ProductsAdapter) GetProductsByText(ctx context.Context, text string) ([]entities.ProductInfo, error) {
	results, err := a.DBConnector.GetFromDatabaseUsingFilter(ctx, ProductsDatabaseName, ProductsCollectionName, bson.D{
		primitive.E{
			Key: "$or", Value: bson.A{
				bson.D{
					primitive.E{
						Key: "brand",
						Value: primitive.Regex{
							Pattern: "" + text + "",
							Options: "",
						},
					}},
				bson.D{
					primitive.E{
						Key: "description",
						Value: primitive.Regex{
							Pattern: "" + text + "",
							Options: "",
						},
					},
				},
			},
		},
	})

	if err != nil {
		a.Log.Error("failed to obtain all databases names: %v", err)
		return []entities.ProductInfo{}, err
	}

	a.Log.Debug("result obtained: %+v", results)
	if len(results) == 0 {
		msg := fmt.Sprintf("no registry for text :%v in database", text)
		a.Log.Error(msg)
		return []entities.ProductInfo{}, errors.New(msg)
	}
	return results, nil
}
