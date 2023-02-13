package main

import (
	"github.com/aluis94/terra-pi-server/middleware"
	"github.com/aluis94/terra-pi-server/router"
	"github.com/urfave/negroni"
)

func main() {
	//gorm
	middleware.InitialMigration()
	//router
	router := router.NewRouter()
	//negroni Middlware
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
