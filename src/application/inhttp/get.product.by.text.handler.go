package inhttp

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/services"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

/*GetProductsByFieldDependencies is a struct definition used for storing dependencies of the GetProdByTextHTTPHandler.
In case you need to add any, add here so as to force dep inyection when invoking.*/
type GetProductsByFieldDependencies struct {
	Svc services.GetProductByTextServiceContract
	Log logger.LogContract
}

/*CreateGetProductByFieldHandlerFunc creates a handler function that retrieves from a query parameter a text to search*/
func CreateGetProductByFieldHandlerFunc(dep GetProductsByFieldDependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		textToSearch := req.URL.Query().Get("text")
		dep.Log.Info("received GET request for text: %v", textToSearch)
		if textToSearch == "" {
			errMsg := fmt.Sprintf("invalid text to search for sent: %v", textToSearch)
			dep.Log.Error(errMsg)
			http.Error(rw, wrapErrAsJSON(errors.New(errMsg)), http.StatusBadRequest)
			return
		}
		//Useful for time outs
		reqContext := req.Context()
		dep.Log.Debug("created context to obtain products")
		products, err := dep.Svc.GetProductsByText(reqContext, textToSearch)
		if err != nil {
			dep.Log.Error("received an error while fetching products by text: %v", err)
			http.Error(rw, wrapErrAsJSON(err), http.StatusInternalServerError)
			return
		}
		dep.Log.Info("mapping product to response")
		dep.Log.Debug("products being mapped: %+v", products)
		response, err := mapProductsToJSONResponse(products)
		if err != nil {
			dep.Log.Error("received an error while generating response: %v", err)
			http.Error(rw, wrapErrAsJSON(err), http.StatusInternalServerError)
			return
		}
		dep.Log.Debug("response to write: %+v", string(response))
		rw.Write(response)
		dep.Log.Info("response has been sent back")
	})

}
