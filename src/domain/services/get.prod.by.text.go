package services

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/ports"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

/*GetProductsByTextServiceDefinition holds the dependencies needed for the service to work*/
type GetProductsByTextServiceDefinition struct {
	Port ports.GetProductsByTextPort
	Log  logger.LogContract
}

/*GetProductsByText stores the service definition to get products via text fields search*/
func (svc GetProductsByTextServiceDefinition) GetProductsByText(reqCtx context.Context, text string) ([]entities.ProductInfo, error) {

	svc.Log.Info("trying to look for a product with text: %v", text)
	if len(text) < 3 {
		msg := fmt.Sprintf("text must be at least of 3  chars long")
		svc.Log.Debug(msg)
		return []entities.ProductInfo{}, errors.New(msg)
	}
	prods, err := svc.Port.GetProductsByText(reqCtx, text)
	if err != nil {
		svc.Log.Error("error while looking for a product by text. text: %v , error: %v", text, err)
		return []entities.ProductInfo{}, err
	}
	svc.Log.Debug("port returned the following product: %+v", prods)
	if len(prods) == 0 {
		msg := fmt.Sprintf("no products found with text: %v", text)
		svc.Log.Debug(msg)
		return []entities.ProductInfo{}, errors.New(msg)
	}

	prodsToReturn := make([]entities.ProductInfo, len(prods))
	for i, prod := range prods {
		if isPalindromeString(text) {
			prodsToReturn[i] = entities.ProductInfo{
				ID:                 prod.ID,
				Title:              prod.Title,
				Description:        prod.Description,
				ImageURL:           prod.ImageURL,
				FullPrice:          prod.FullPrice,
				FinalPrice:         prod.FullPrice * DiscountByPalindrome,
				PriceModifications: -DiscountByPalindrome,
			}
			continue
		}
		prodsToReturn[i] = entities.ProductInfo{
			ID:                 prod.ID,
			Title:              prod.Title,
			Description:        prod.Description,
			ImageURL:           prod.ImageURL,
			FullPrice:          prod.FullPrice,
			FinalPrice:         prod.FullPrice,
			PriceModifications: 0,
		}
	}

	return prodsToReturn, err
}
