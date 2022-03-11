package services

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/ports"
)

func GetProductByIdServDef(getProdByIdPort ports.GetProductByIdPort) func(id uint, ctx context.Context) (entities.ProductInfo, error) {
	return func(id uint, ctx context.Context) (entities.ProductInfo, error) {
		prod, err := getProdByIdPort.GetProductsById(id, ctx)
		return prod, err
	}

}
