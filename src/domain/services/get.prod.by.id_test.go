package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

func TestGetAProduct(t *testing.T) {
	testCase := getProdByIDTestCase{

		testName: "retrieve one product, not a palindrome",
		id:       123,
		existingProductsInPortMock: map[int]entities.ProductInfo{
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

func TestGetProductWithPalindromeId(t *testing.T) {
	testCase := getProdByIDTestCase{

		testName: "retrieve a product with palindrome id",
		id:       181,
		existingProductsInPortMock: map[int]entities.ProductInfo{
			181: {
				Id:          181,
				Title:       "a palindromic(?) product",
				FullPrice:   1000,
				Description: "palindrome",
			},
		},
		expectedProd: entities.ProductInfo{
			Id:                 181,
			Title:              "a palindromic(?) product",
			FullPrice:          1000,
			FinalPrice:         500,
			PriceModifications: -0.5,
			Description:        "palindrome",
		},
		expectedErr: "",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}

func TestNoProductFound(t *testing.T) {
	testCase := getProdByIDTestCase{

		testName: "retrieve no product as id doesn't match any",
		id:       55,
		existingProductsInPortMock: map[int]entities.ProductInfo{
			181: {
				Id:          181,
				Title:       "a palindromic(?) product",
				FullPrice:   1000,
				Description: "palindrome",
			},
		},
		expectedProd: entities.ProductInfo{},
		expectedErr:  "",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}

// End test cases

type getProdByIDTestCase struct {
	testName                   string
	id                         int
	existingProductsInPortMock map[int]entities.ProductInfo
	errorPortInMock            error
	expectedProd               entities.ProductInfo
	expectedErr                string
}

func (testCase getProdByIDTestCase) testAndAssert(t *testing.T) {
	t.Logf("testing function")

	/**---------------------- FUNCTION UNDER TEST -----------------------**/
	/*Dependencies*/
	mockedPort := mockPort{
		products: testCase.existingProductsInPortMock,
		err:      testCase.errorPortInMock,
	}
	loggerFactory := logger.LogFactory{LogLevel: "INFO"}
	log := loggerFactory.CreateLog("")
	/*END Dependencies*/
	serviceDef := GetProductByIdServDef(mockedPort, log)
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
	products map[int]entities.ProductInfo
	err      error
}

func (mock mockPort) GetProductsById(id int, ctx context.Context) (entities.ProductInfo, error) {
	return mock.products[id], mock.err
}
