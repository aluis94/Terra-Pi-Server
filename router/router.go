package router

import (
	gmux "github.com/gorilla/mux"
)

//NewRouter Method
func NewRouter() *gmux.Router {
	router := gmux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
