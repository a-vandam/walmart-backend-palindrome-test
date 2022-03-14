package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/entities"
	"gitlab.com/a.vandam/product-search-challenge/src/domain/services"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

func CreateGetProdByIdHandlerFunc(svc services.GetProductByIdServiceContract, log logger.LogContract) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		idInPath := req.Context().Value(ProductIdCtxKey{})
		log.Info("received GET request for id: %v", idInPath)
		pathId := parseInterfToInt(idInPath)
		if pathId == DefaultErrorInt || pathId == 0 || idInPath == nil {
			errMsg := fmt.Sprintf("invalid product id path parameter sent: received: %v", idInPath)
			log.Error(errMsg)
			http.Error(rw, wrapErrAsJson(errors.New(errMsg)), http.StatusBadRequest)
			return
		}

		//Useful for time outs
		reqContext := req.Context()
		log.Debug("created context to obtain products")
		product, err := svc.GetProductsById(pathId, reqContext)
		if err != nil {
			log.Error("received an error while fetching product by id: %v", err)
			http.Error(rw, wrapErrAsJson(err), http.StatusInternalServerError)
			return
		}
		log.Info("mapping product to response")
		log.Debug("product being mapped: %+v", product)
		response, err := mapProductToJsonResponse(&product)
		if err != nil {
			log.Error("received an error while generating response: %v", err)
			http.Error(rw, wrapErrAsJson(err), http.StatusInternalServerError)
			return
		}
		log.Debug("response to write: %+v", response)
		rw.Write(response)
		log.Info("response has been sent back")
	})

}

func mapProductToJsonResponse(prod *entities.ProductInfo) ([]byte, error) {
	responseDTO := embeddingJsonResponse{
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

type embeddingJsonResponse struct {
	ErrMsg    string                 `json:"error"`
	Resources getProductJsonResponse `json:"resources"`
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
