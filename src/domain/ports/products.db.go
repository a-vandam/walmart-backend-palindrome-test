package ports

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

type GetProductByIdPort interface {
	GetProductById(id int, ctx context.Context) (entities.ProductInfo, error)
}
