package inhttp

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

type ProductIdPathParamCtxKey struct {
}

const DefaultErrorInt int = 0

func parseInterfaceToInt(element interface{}) int {

	valueAsInt, parseableToInt := element.(int)
	if parseableToInt {
		return valueAsInt
	}
	valueAsString, parseable := element.(string)

	if !parseable {
		return DefaultErrorInt
	}

	valueAsInt, err := strconv.Atoi(valueAsString)
	if err != nil {
		return 0
	}
	defer func() {
		parseConvErr := recover()
		if parseConvErr != nil {
			valueAsInt = DefaultErrorInt
		}
	}()
	return valueAsInt
}

func parseIdPathParamToInt(element interface{}) int {
	elementAsInts, parseableToArray := element.([]int)
	if !parseableToArray {
		return DefaultErrorInt
	}
	elementAsInt := elementAsInts[0]
	defer func() {
		parseConvErr := recover()
		if parseConvErr != nil {
			elementAsInt = DefaultErrorInt
		}
	}()
	return elementAsInt
}

const ErrorResponseBody string = `{"error":"%v", "resources":{}}`

func wrapErrAsJson(err error) string {
	return fmt.Sprintf(ErrorResponseBody, err.Error())
}

type embeddingOneResourceJsonResponse struct {
	ErrMsg    string                 `json:"error"`
	Resources getProductJsonResponse `json:"resources"`
}

type embeddingMultipleResourcesJsonResponse struct {
	ErrMsg    string                   `json:"error"`
	Resources []getProductJsonResponse `json:"resources"`
}
type getProductJsonResponse struct {
	Id                 int     `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	ImageURL           string  `json:"imageURL"`
	FullPrice          float32 `json:"fullPrice"`
	FinalPrice         float32 `json:"finalPrice"`
	PriceModifications float32 `json:"priceModifications"`
}

func mapProductToJsonResponse(prod *entities.ProductInfo) ([]byte, error) {
	responseDTO := embeddingOneResourceJsonResponse{
		ErrMsg: "",
		Resources: getProductJsonResponse{
			Id:                 prod.Id,
			Title:              prod.Title,
			Description:        prod.Description,
			ImageURL:           prod.ImageURL,
			FullPrice:          prod.FullPrice,
			FinalPrice:         prod.FinalPrice,
			PriceModifications: prod.PriceModifications,
		},
	}
	responseBody, err := json.Marshal(responseDTO)

	return responseBody, err
}
func mapProductsToJsonResponse(prods []entities.ProductInfo) ([]byte, error) {
	resources := make([]getProductJsonResponse, len(prods))
	for i, prod := range prods {
		resources[i] = getProductJsonResponse{
			Id:                 prod.Id,
			Title:              prod.Title,
			Description:        prod.Description,
			ImageURL:           prod.ImageURL,
			FullPrice:          prod.FullPrice,
			FinalPrice:         prod.FinalPrice,
			PriceModifications: prod.PriceModifications,
		}
	}
	response := embeddingMultipleResourcesJsonResponse{ErrMsg: "", Resources: resources}
	responseBody, err := json.Marshal(response)

	return responseBody, err
}
