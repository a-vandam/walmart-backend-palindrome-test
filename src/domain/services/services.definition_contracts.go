package services

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

type GetProductByIdServiceContract interface {
	GetProductsByIdService(id int32, ctx context.Context) ([]entities.ProductInfo, error)
}
