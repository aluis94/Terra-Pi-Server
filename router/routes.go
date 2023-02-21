package router

import (
	"net/http"

	"github.com/aluis94/terra-pi-server/middleware"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slice
type Routes []Route

var routes = Routes{
	Route{
		"Home",
		"GET",
		"/",
		middleware.Home,
	},
	//add the rest of the routes here
}
