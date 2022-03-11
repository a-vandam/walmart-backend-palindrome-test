package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

func TestGetAProduct(t *testing.T) {
	testCase := getProdByIDTestCase{

		testName: "retrieve one product, not a palindrome",
		id:       123,
		existingProductsInPortMock: map[uint]entities.ProductInfo{
			123: {
				Id:                 123,
				Title:              "a random product",
				FullPrice:          1000,
				FinalPrice:         1000,
				PriceModifications: 0.0,
			},
		},
		expectedProd: entities.ProductInfo{
			Id:                 123,
			Title:              "a random product",
			FullPrice:          1000,
			FinalPrice:         1000,
			PriceModifications: 0.0,
		},
		expectedErr: "",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}

type getProdByIDTestCase struct {
	testName                   string
	id                         uint
	existingProductsInPortMock map[uint]entities.ProductInfo
	errorPortInMock            error
	expectedProd               entities.ProductInfo
	expectedErr                string
}

func (testCase getProdByIDTestCase) testAndAssert(t *testing.T) {
	t.Logf("testing function")

	/**---------------------- FUNCTION UNDER TEST -----------------------**/
	mockedPort := mockPort{
		products: testCase.existingProductsInPortMock,
		err:      testCase.errorPortInMock,
	}
	serviceDef := GetProductByIdServDef(mockedPort)
	testCtx := context.Background()
	product, err := serviceDef(testCase.id, testCtx)
	/**---------------------- END FUNCTION UNDER TEST -----------------------**/

	if !assert.Equal(t, testCase.expectedProd, product, "difference in value expected (%v) and obtained (%v)", testCase.expectedProd, product) {
		t.Fail()
	}
	if testCase.expectedErr != "" && err == nil {
		t.Logf("test failed as the function did not return an expected error: %v vs %v", err, testCase.expectedErr)
		t.FailNow()
	}
	if testCase.expectedErr == "" && err != nil {
		t.Logf("test failed as the function returned an error when it shouldn't: %v", err)
		t.FailNow()
	}
	if testCase.expectedErr != "" && err != nil {
		//comparing errors
		if assert.EqualErrorf(t, err, testCase.expectedErr, "function returned an unexpected error: expected: %v vs found: %v", testCase.expectedErr, err) {
			t.FailNow()
		}
	}

	t.Logf("OK!!!! - test case:  %v  - OK!!!!", testCase.testName)
}

/*Mocking*/
type mockPort struct {
	products map[uint]entities.ProductInfo
	err      error
}

func (mock mockPort) GetProductsById(id uint, ctx context.Context) (entities.ProductInfo, error) {
	return mock.products[id], mock.err
}
