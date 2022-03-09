package services

import (
	"testing"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

func TestGetProductByIdServDef(t *testing.T) {

	GetProductByIdServDef()
}

type GetProdByIDTestCase struct {
	testName     string
	id           uint
	expectedProd entities.ProductInfo
	expectedErr  string
}
