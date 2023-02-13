package router

import (
	"net/http"

	"github.com/aluis94/terra-pi-server/middleware"
)

//Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes slice
type Routes []Route

var routes = Routes{
	Route{
		"Home",
		"GET",
		"/",
		middleware.Home,
	},
	//Position
	Route{
		"AddPosition",
		"POST",
		"/position/add",
		middleware.AddPosition,
	},
	Route{
		"AddPosition",
		"OPTIONS",
		"/position/add",
		middleware.AddPosition,
	},
	Route{
		"EditPosition",
		"PUT",
		"/position/edit/{id}",
		middleware.EditPosition,
	},
	Route{
		"EditPosition",
		"OPTIONS",
		"/position/edit/{id}",
		middleware.EditPosition,
	},
	Route{
		"DeletePosition",
		"DELETE",
		"/position/delete/{id}",
		middleware.DeletePosition,
	},
	Route{
		"DeletePosition",
		"OPTIONS",
		"/position/delete/{id}",
		middleware.DeletePosition,
	},
	{
		"ViewPositions",
		"GET",
		"/position/view-all",
		middleware.ViewPositions,
	},
	Route{
		"ViewPosition",
		"GET",
		"/position/view/{id}",
		middleware.ViewPosition,
	},
	//Companies
	{
		"AddCompany",
		"POST",
		"/company/add",
		middleware.AddCompany,
	},
	Route{
		"DeleteCompany",
		"OPTIONS",
		"/company/delete/{id}",
		middleware.DeleteCompany,
	},
	Route{
		"DeleteCompany",
		"DELETE",
		"/company/delete/{id}",
		middleware.DeleteCompany,
	},
	Route{
		"ViewCompany",
		"GET",
		"/company/view/{id}",
		middleware.ViewCompany,
	},
	Route{
		"ViewCompanies",
		"GET",
		"/company/view-all",
		middleware.ViewCompanies,
	},
	Route{
		"EditCompany",
		"PUT",
		"/company/edit/{id}",
		middleware.EditCompany,
	},
	Route{
		"EditCompany",
		"OPTIONS",
		"/company/edit/{id}",
		middleware.EditCompany,
	},
}
