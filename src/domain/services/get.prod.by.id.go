package services

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/ports"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

const DiscountByPalindromeIdCoefficient float32 = 0.5

func GetProductByIdServDef(getProdByIdPort ports.GetProductByIdPort, log logger.LogContract) func(id int, ctx context.Context) (entities.ProductInfo, error) {
	return func(id int, ctx context.Context) (entities.ProductInfo, error) {
		log.Info("trying to look for a product with id: %v", id)
		prod, err := getProdByIdPort.GetProductsById(id, ctx)
		if err != nil {
			log.Error("error while looking for a product by id. id: %v , error: %v", id, err)
			return entities.ProductInfo{}, err
		}
		if prod.Title == "" {
			log.Debug("no products found with id: %v", id)
			return prod, nil
		}

		if isPalindromeInt(prod.Id) {
			applyDiscount(&prod, DiscountByPalindromeIdCoefficient)
		}
		return prod, err
	}
}
func applyDiscount(product *entities.ProductInfo, discount float32) {
	product.FinalPrice = product.FullPrice * discount
	product.PriceModifications = -1 * discount
}
