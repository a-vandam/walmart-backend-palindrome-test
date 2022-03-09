package services

import "context"

type GetProductByIdServiceContract interface {
	GetProductByIdService(id int32, ctx context.Context)
}
