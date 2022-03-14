package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

func TestGetExistingProdWithIdInPath(t *testing.T) {
	testCase := getProdByIDTestReq{
		testName: "request with path param sent",
		path:     "http://svctest/products/23",
		id:       23,
		bodyFile: "",
		verb:     "GET",
		svcProdsResponse: entities.ProductInfo{
			Id:                 23,
			Title:              "a product",
			Description:        "i'm a merchandise",
			ImageURL:           "http://blablabla",
			FullPrice:          100,
			FinalPrice:         80,
			PriceModifications: -0.2,
		},
		expectedCode:     200,
		expectedJsonResp: "json_examples/get.product.by.id_ok.response.json",
	}

	t.Run(testCase.testName, testCase.testAndAssert)
}

type getProdByIDTestReq struct {
	testName         string
	path             string
	id               int
	verb             string
	bodyFile         string
	svcProdsResponse entities.ProductInfo
	svcErrResponse   error
	expectedCode     int
	expectedJsonResp string
}

func (testCase getProdByIDTestReq) testAndAssert(t *testing.T) {
	t.Logf("testing function")

	/**---------------------- FUNCTION UNDER TEST -----------------------**/
	/*Dependencies*/
	mockedSvc := getProdByIdSvcMock{
		product: testCase.svcProdsResponse,
		svcErr:  testCase.svcErrResponse,
	}
	loggerFactory := logger.LogFactory{LogLevel: "INFO"}
	log := loggerFactory.CreateLog("")
	/*END Dependencies*/
	// Test function //
	handlerFunc := CreateGetProdByIdHandlerFunc(mockedSvc, log)

	/*Create test request and server*/
	var req *http.Request
	requestContext := context.WithValue(context.Background(), ProductIdCtxKey{}, testCase.id)

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
	req = req.WithContext(requestContext)

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
	var expectedBodyAsMap, receivedBodyAsMap interface{}
	json.Unmarshal(expectedBody, &expectedBodyAsMap)
	receivedBody := responseWriter.Body.Bytes()
	json.Unmarshal(receivedBody, &receivedBodyAsMap)

	if !assert.Equal(t, expectedBodyAsMap, receivedBodyAsMap, "difference in body expected (%v) and obtained (%v)", expectedBodyAsMap, receivedBodyAsMap) {
		t.FailNow()
		return
	}

	t.Logf("OK!!!! - test case:  %v  - OK!!!!", testCase.testName)
}

type getProdByIdSvcMock struct {
	product entities.ProductInfo
	svcErr  error
}

func (mock getProdByIdSvcMock) GetProductsById(id int, ctx context.Context) (entities.ProductInfo, error) {
	return mock.product, mock.svcErr
}
