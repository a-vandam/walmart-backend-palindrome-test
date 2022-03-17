package inhttp

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
)

/*Key used for retrieving path params from the req's context.*/
type productIDPathParamCtxKey struct {
}

const defaultParseErrorInt int = 0

func parseInterfaceToInt(element interface{}) int {

	valueAsInt, parseableToInt := element.(int)
	if parseableToInt {
		return valueAsInt
	}
	valueAsString, parseable := element.(string)

	if !parseable {
		return defaultParseErrorInt
	}

	valueAsInt, err := strconv.Atoi(valueAsString)
	if err != nil {
		return 0
	}
	defer func() {
		parseConvErr := recover()
		if parseConvErr != nil {
			valueAsInt = defaultParseErrorInt
		}
	}()
	return valueAsInt
}

func parseIDPathParamToInt(element interface{}) int {
	elementAsInts, parseableToArray := element.([]int)
	if !parseableToArray {
		return defaultParseErrorInt
	}
	elementAsInt := elementAsInts[0]
	defer func() {
		parseConvErr := recover()
		if parseConvErr != nil {
			elementAsInt = defaultParseErrorInt
		}
	}()
	return elementAsInt
}

const errorResponseBody string = `{"error":"%v", "resources":{}}`

func wrapErrAsJSON(err error) string {
	return fmt.Sprintf(errorResponseBody, err.Error())
}

type embeddingOneResourceJSONResponse struct {
	ErrMsg    string                 `json:"error"`
	Resources getProductJSONResponse `json:"resources"`
}

type embeddingMultipleResourcesJSONResponse struct {
	ErrMsg    string                   `json:"error"`
	Resources []getProductJSONResponse `json:"resources"`
}
type getProductJSONResponse struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	ImageURL           string  `json:"imageURL"`
	FullPrice          float32 `json:"fullPrice"`
	FinalPrice         float32 `json:"finalPrice"`
	PriceModifications float32 `json:"priceModifications"`
}

func mapProductToJSONResponse(prod *entities.ProductInfo) ([]byte, error) {
	responseDTO := embeddingOneResourceJSONResponse{
		ErrMsg: "",
		Resources: getProductJSONResponse{
			ID:                 prod.ID,
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
func mapProductsToJSONResponse(prods []entities.ProductInfo) ([]byte, error) {
	resources := make([]getProductJSONResponse, len(prods))
	for i, prod := range prods {
		resources[i] = getProductJSONResponse{
			ID:                 prod.ID,
			Title:              prod.Title,
			Description:        prod.Description,
			ImageURL:           prod.ImageURL,
			FullPrice:          prod.FullPrice,
			FinalPrice:         prod.FinalPrice,
			PriceModifications: prod.PriceModifications,
		}
	}
	response := embeddingMultipleResourcesJSONResponse{ErrMsg: "", Resources: resources}
	responseBody, err := json.Marshal(response)

	return responseBody, err
}
