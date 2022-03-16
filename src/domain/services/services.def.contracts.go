package services

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

type GetProductByIdServiceContract interface {
	GetProductById(id int, ctx context.Context) (entities.ProductInfo, error)
}
