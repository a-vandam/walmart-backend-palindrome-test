package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

func TestGetProductsByText(t *testing.T) {
	testCase := getProdsByTextTestCase{
		textToSearch: "example",
		testName:     "retrieve products not palindrome",
		returnedProdsInPortMock: []entities.ProductInfo{
			{
				ID:                 9999,
				Title:              "a example random product",
				FullPrice:          1000,
				FinalPrice:         1000,
				PriceModifications: 0.0,
			},
			{
				ID:                 123,
				Title:              "a random product",
				Description:        "example desc",
				FullPrice:          1000,
				FinalPrice:         1000,
				PriceModifications: 0.0,
			},
		},
		expectedProd: []entities.ProductInfo{
			{
				ID:                 9999,
				Title:              "a example random product",
				FullPrice:          1000,
				FinalPrice:         1000,
				PriceModifications: 0.0,
			},
			{
				ID:                 123,
				Title:              "a random product",
				Description:        "example desc",
				FullPrice:          1000,
				FinalPrice:         1000,
				PriceModifications: 0.0,
			},
		},
		expectedErr: "",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}

func TestShorterThan3CharsSearchProductsByText(t *testing.T) {
	testCase := getProdsByTextTestCase{
		textToSearch: "ab",
		testName:     "error when searching for less than 3 word",
		returnedProdsInPortMock: []entities.ProductInfo{
			{
				ID:                 9999,
				Title:              "a ab example random product",
				FullPrice:          1000,
				FinalPrice:         1000,
				PriceModifications: 0.0,
			},
			{
				ID:                 123,
				Title:              "a random product",
				Description:        "example ab desc",
				FullPrice:          1000,
				FinalPrice:         1000,
				PriceModifications: 0.0,
			},
		},
		expectedProd: []entities.ProductInfo{},
		expectedErr:  "text must be at least of 3  chars long",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}
func TestGetProductsWithPalindromeTexts(t *testing.T) {
	testCase := getProdsByTextTestCase{

		testName:     "retrieve a product with palindrome text",
		textToSearch: "abba",
		returnedProdsInPortMock: []entities.ProductInfo{
			{
				ID:                 9999,
				Title:              "a example abba random product",
				FullPrice:          1000,
				FinalPrice:         0,
				PriceModifications: 0.0,
			},
			{
				ID:                 123,
				Title:              "a random product",
				Description:        "example abba desc",
				FullPrice:          2000,
				FinalPrice:         0,
				PriceModifications: 0.0,
			},
		},
		expectedProd: []entities.ProductInfo{
			{
				ID:                 9999,
				Title:              "a example abba random product",
				FullPrice:          1000,
				FinalPrice:         500,
				PriceModifications: -0.5,
			},
			{
				ID:                 123,
				Title:              "a random product",
				Description:        "example abba desc",
				FullPrice:          2000,
				FinalPrice:         1000,
				PriceModifications: -0.5,
			},
		},
		expectedErr: "",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}

func TestNoProductsFound(t *testing.T) {
	testCase := getProdsByTextTestCase{

		testName:                "retrieve no product as text doesn't match any",
		textToSearch:            "asdasidwaji",
		returnedProdsInPortMock: []entities.ProductInfo{},
		expectedProd:            []entities.ProductInfo{},
		errorPortInMock:         nil,
		expectedErr:             "no products found with text: asdasidwaji",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}

// End test cases

type getProdsByTextTestCase struct {
	testName                string
	textToSearch            string
	returnedProdsInPortMock []entities.ProductInfo
	errorPortInMock         error
	expectedProd            []entities.ProductInfo
	expectedErr             string
}

func (testCase getProdsByTextTestCase) testAndAssert(t *testing.T) {
	t.Logf("testing function")

	/**---------------------- FUNCTION UNDER TEST -----------------------**/
	/*Dependencies*/
	mockedPort := mockGetProductByTextPort{
		products: testCase.returnedProdsInPortMock,
		err:      testCase.errorPortInMock,
	}
	loggerFactory := logger.LogFactory{LogLevel: "INFO"}
	svc := GetProductsByTextServiceDefinition{
		Port: mockedPort,
		Log:  loggerFactory.CreateLog(""),
	}
	/*END Dependencies*/
	testCtx := context.Background()
	product, err := svc.GetProductsByText(testCtx, testCase.textToSearch)
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
		if !assert.EqualErrorf(t, err, testCase.expectedErr, "function returned an unexpected error: expected: %v vs found: %v", testCase.expectedErr, err) {
			t.FailNow()
		}
	}

	t.Logf("OK!!!! - test case:  %v  - OK!!!!", testCase.testName)
}

/*Mocking port*/
type mockGetProductByTextPort struct {
	products []entities.ProductInfo
	err      error
}

/*Mocking port's functions*/
func (mock mockGetProductByTextPort) GetProductsByText(ctx context.Context, text string) ([]entities.ProductInfo, error) {
	return mock.products, mock.err
}
