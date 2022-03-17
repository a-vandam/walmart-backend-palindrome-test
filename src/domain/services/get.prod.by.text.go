package services

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/ports"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

type GetProductsByTextServiceDefinition struct {
	Port ports.GetProductsByTextPort
	Log  logger.LogContract
}

func (svc GetProductsByTextServiceDefinition) GetProductsByText(text string, reqCtx context.Context) ([]entities.ProductInfo, error) {

	svc.Log.Info("trying to look for a product with text: %v", text)
	if len(text) < 3 {
		msg := fmt.Sprintf("text must be at least of 3  chars long")
		svc.Log.Debug(msg)
		return []entities.ProductInfo{}, errors.New(msg)
	}
	prods, err := svc.Port.GetProductsByText(text, reqCtx)
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
				Id:                 prod.Id,
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
			Id:                 prod.Id,
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
