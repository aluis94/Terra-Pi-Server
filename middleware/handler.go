package middleware

import (
	"fmt"
	"net/http"
)

// Home handlefunc
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Home Endpoint Hit")

}

//add handler functions here
