package inhttp

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"gitlab.com/a.vandam/product-search-challenge/src/logger"
)

/*ProductsRouter is a router struct that stores both dependencies needed for the URL Path router to work, and the routes which will then be router by it using regexp.
New routes should be added with the Add Router method*/
type ProductsRouter struct {
	routes []route
	Log    logger.LogContract
}

/*Serve should be invoked by the main server in combination with ListenAndServe. It generates a middleware that allows routing using path parameters and regexp. */
func (router ProductsRouter) Serve(w http.ResponseWriter, r *http.Request) {

	var allow []string
	router.Log.Info("routing request with method + path: %v: %v", r.Method, r.URL.Path)
	for _, route := range router.routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), productIDPathParamCtxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

/*RegisteredRoutes allows to retrieve the router's registered routes. Useful for debugging purposes.*/
func (router ProductsRouter) RegisteredRoutes() []string {
	paths := make([]string, len(router.routes))
	for _, route := range router.routes {
		paths = append(paths, route.method+" - path - "+route.regex.String())
	}
	return paths
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

/*AddRoute adds a Route struct to the Router. This should be invoked when adding routess with their handlers.*/
func (router ProductsRouter) AddRoute(verb string, path string, handlerFunc http.HandlerFunc) ProductsRouter {
	router.Log.Debug("registering route: verb: %v,subpath: %v", verb, path)
	router.routes = append(router.routes, newRoute(verb, path, handlerFunc))
	return router
}
