package services

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/ports"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

const DiscountByPalindromeIdCoefficient float32 = 0.5

type GetProductByIdServiceDefinition struct {
	Port ports.GetProductByIdPort
	Log  logger.LogContract
}

func (svc GetProductByIdServiceDefinition) GetProductById(id int, ctx context.Context) (entities.ProductInfo, error) {

	svc.Log.Info("trying to look for a product with id: %v", id)
	prod, err := svc.Port.GetProductById(id, ctx)
	if err != nil {
		svc.Log.Error("error while looking for a product by id. id: %v , error: %v", id, err)
		return entities.ProductInfo{}, err
	}
	svc.Log.Debug("port returned the following product: %+v", prod)
	if prod.Title == "" {
		msg := fmt.Sprintf("no products found with id: %v", id)
		svc.Log.Debug(msg)
		return entities.ProductInfo{}, errors.New(msg)
	}

	if isPalindromeInt(prod.Id) {
		svc.Log.Info("applying discount to product")
		applyDiscount(&prod, DiscountByPalindromeIdCoefficient)
	}
	return prod, err
}
func applyDiscount(product *entities.ProductInfo, discount float32) {
	product.FinalPrice = product.FullPrice * discount
	product.PriceModifications = -1 * discount
}
