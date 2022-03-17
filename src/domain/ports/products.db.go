package ports

import (
	"context"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

/*GetProductByIDPort is the port definition to get a product via id*/
type GetProductByIDPort interface {
	GetProductByID(ctx context.Context, id int) (entities.ProductInfo, error)
}

/*GetProductsByTextPort is the Port definition to get a product by looking for text within it*/
type GetProductsByTextPort interface {
	GetProductsByText(ctx context.Context, text string) ([]entities.ProductInfo, error)
}
