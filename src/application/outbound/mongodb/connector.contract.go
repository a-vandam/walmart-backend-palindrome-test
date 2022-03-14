package mongodb

type GetAllProductsContract interface {
	GetAll(query string)
}

type GetProductByIdContract interface {
	GetProductById(id uint32)
}
