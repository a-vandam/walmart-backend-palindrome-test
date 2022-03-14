package ports

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

type GetProductByIdPort interface {
	GetProductsById(id int, ctx context.Context) (entities.ProductInfo, error)
}
