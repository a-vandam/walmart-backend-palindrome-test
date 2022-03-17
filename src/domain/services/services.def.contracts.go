package services

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

/*GetProductByIDServiceContract is the contract that service must respect to be injected in the handler*/
type GetProductByIDServiceContract interface {
	GetProductByID(ctx context.Context, id int) (entities.ProductInfo, error)
}

/*GetProductByTextServiceContract is the contract that service must respect to be injected in the handler*/
type GetProductByTextServiceContract interface {
	GetProductsByText(ctx context.Context, string string) ([]entities.ProductInfo, error)
}
