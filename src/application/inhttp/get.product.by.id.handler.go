package inhttp

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gitlab.com/a.vandam/product-search-challenge/src/domain/services"
	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

/*GetProductByIDHandlerDependencies is an struct definition for dependencies of the GetProdByIdHTTPHandler.
In case you need to add, add here so as to force dep inyection when invoking.*/
type GetProductByIDHandlerDependencies struct {
	Svc services.GetProductByIDServiceContract
	Log logger.LogContract
}

/*CreateGetProdByIDHandlerFunc creates a http handler fucntion that retrieves from the reqquest's context a path parameter, if any. IT's existence should be guaranteed
by a middleware / regex router.*/
func CreateGetProdByIDHandlerFunc(dep GetProductByIDHandlerDependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		pathParamReceived := req.Context().Value(productIDPathParamCtxKey{}).([]string)
		idInPath, _ := strconv.Atoi(pathParamReceived[0])
		dep.Log.Info("received GET request for id: %v", idInPath)
		if idInPath == defaultParseErrorInt || idInPath == 0 {
			errMsg := fmt.Sprintf("invalid product id path parameter sent: received: %v", idInPath)
			dep.Log.Error(errMsg)
			http.Error(rw, wrapErrAsJSON(errors.New(errMsg)), http.StatusBadRequest)
			return
		}

		//Useful for time outs
		reqContext := req.Context()
		dep.Log.Debug("created context to obtain products")
		product, err := dep.Svc.GetProductByID(reqContext, idInPath)
		if err != nil {
			dep.Log.Error("received an error while fetching product by id: %v", err)
			http.Error(rw, wrapErrAsJSON(err), http.StatusInternalServerError)
			return
		}
		dep.Log.Info("mapping product to response")
		dep.Log.Debug("product being mapped: %+v", product)
		response, err := mapProductToJSONResponse(&product)
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
