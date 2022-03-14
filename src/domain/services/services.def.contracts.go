package services

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

type GetProductByIdServiceContract interface {
	GetProductsById(id int, ctx context.Context) (entities.ProductInfo, error)
}
