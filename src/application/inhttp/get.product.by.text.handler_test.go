package inhttp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

func TestGetExistingProdByText(t *testing.T) {
	testCase := getProdByTextTestReq{
		testName: "request with path param sent",
		path:     "http://svctest/products/search?text=example",
		bodyFile: "",
		verb:     "GET",
		svcProdsResponse: []entities.ProductInfo{
			{
				Id:                 99,
				Title:              "example title",
				Description:        "i'm a merchandise",
				ImageURL:           "http://blablabla",
				FullPrice:          100,
				FinalPrice:         100,
				PriceModifications: 0,
			},
			{
				Id:                 150,
				Title:              "title",
				Description:        "i'm an example merchandise",
				ImageURL:           "http://blablabla",
				FullPrice:          990,
				FinalPrice:         990,
				PriceModifications: 0,
			},
		},
		expectedCode:     200,
		expectedJsonResp: "json_examples/get.product.by.text_ok.response.json",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}
func TestRequestToGetNonExistentProductByTitle(t *testing.T) {
	testCase := getProdByTextTestReq{
		testName:         "request with path param sent but no prod",
		path:             "http://svctest/products/search?text=example",
		bodyFile:         "",
		verb:             "GET",
		svcProdsResponse: []entities.ProductInfo{{}},
		svcErrResponse:   errors.New("no registry for text : example in database"),
		expectedCode:     500,
		expectedJsonResp: "json_examples/get.product.by.text_no.prod.found.json",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}

type getProdByTextTestReq struct {
	testName         string
	path             string
	verb             string
	bodyFile         string
	svcProdsResponse []entities.ProductInfo
	svcErrResponse   error
	expectedCode     int
	expectedJsonResp string
}

func (testCase getProdByTextTestReq) testAndAssert(t *testing.T) {
	t.Logf("testing function")

	/**---------------------- FUNCTION UNDER TEST -----------------------**/
	/*Dependencies*/
	mockedSvc := getProdByTextSvcMock{
		product: testCase.svcProdsResponse,
		svcErr:  testCase.svcErrResponse,
	}
	loggerFactory := logger.LogFactory{LogLevel: "DEBUG"}
	log := loggerFactory.CreateLog("")
	dependencies := GetProductsByField{mockedSvc, log}
	/*END Dependencies*/
	// Test function //
	handlerFunc := CreateGetProductByFieldHandlerFunc(dependencies)

	/*Create test request and server*/
	var req *http.Request

	if testCase.bodyFile == "" {
		req = httptest.NewRequest(testCase.verb, testCase.path, nil)
	} else {
		bodyToSend, fileReadErr := os.ReadFile(testCase.bodyFile)
		if fileReadErr != nil {
			t.Logf("error while opening json request file: %v", fileReadErr)
			t.FailNow()
			return
		}
		reader := strings.NewReader(string(bodyToSend))
		req = httptest.NewRequest(testCase.verb, testCase.path, reader)

	}
	responseWriter := httptest.NewRecorder()
	handlerFunc(responseWriter, req)
	/*END Create test request and server*/
	/**---------------------- END FUNCTION UNDER TEST -----------------------**/
	receivedCode := responseWriter.Result().StatusCode
	if !assert.Equal(t, testCase.expectedCode, receivedCode, "difference in http code expected (%v) and obtained (%v)", testCase.expectedCode, receivedCode) {
		t.FailNow()
		return
	}
	expectedBody, bodyNotFoundErr := os.ReadFile(testCase.expectedJsonResp)
	if bodyNotFoundErr != nil {
		t.Logf("json file that stores the expected body has not been found: %v", bodyNotFoundErr)
		t.FailNow()
		return
	}
	var expectedBodyAsMap, receivedBodyAsMap embeddingMultipleResourcesJsonResponse
	json.Unmarshal(expectedBody, &expectedBodyAsMap)
	receivedBody := responseWriter.Body.Bytes()
	json.Unmarshal(receivedBody, &receivedBodyAsMap)

	if !assert.Equal(t, expectedBodyAsMap, receivedBodyAsMap, "difference in body expected (%v) and obtained (%v)", expectedBodyAsMap, receivedBodyAsMap) {
		t.FailNow()
		return
	}

	t.Logf("OK!!!! - test case:  %v  - OK!!!!", testCase.testName)
}

type getProdByTextSvcMock struct {
	product []entities.ProductInfo
	svcErr  error
}

func (mock getProdByTextSvcMock) GetProductsByText(text string, ctx context.Context) ([]entities.ProductInfo, error) {
	return mock.product, mock.svcErr
}
