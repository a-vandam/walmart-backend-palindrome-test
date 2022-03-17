package inhttp

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/services"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

type GetProductByField struct {
	Svc services.GetProductByTextServiceContract
	Log logger.LogContract
}

func CreateGetProductByFieldHandlerFunc(dep GetProductByField) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		textToSearch := req.URL.Query().Get("text")
		dep.Log.Info("received GET request for text: %v", textToSearch)
		if textToSearch == "" {
			errMsg := fmt.Sprintf("invalid text to search for sent: %v", textToSearch)
			dep.Log.Error(errMsg)
			http.Error(rw, wrapErrAsJson(errors.New(errMsg)), http.StatusBadRequest)
			return
		}
		//Useful for time outs
		reqContext := req.Context()
		dep.Log.Debug("created context to obtain products")
		products, err := dep.Svc.GetProductsByText(textToSearch, reqContext)
		if err != nil {
			dep.Log.Error("received an error while fetching products by text: %v", err)
			http.Error(rw, wrapErrAsJson(err), http.StatusInternalServerError)
			return
		}
		dep.Log.Info("mapping product to response")
		dep.Log.Debug("products being mapped: %+v", products)
		response, err := mapProductsToJsonResponse(products)
		if err != nil {
			dep.Log.Error("received an error while generating response: %v", err)
			http.Error(rw, wrapErrAsJson(err), http.StatusInternalServerError)
			return
		}
		dep.Log.Debug("response to write: %+v", string(response))
		rw.Write(response)
		dep.Log.Info("response has been sent back")
	})

}
