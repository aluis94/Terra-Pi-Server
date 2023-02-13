package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aluis94/terra-pi-server/models"
	"github.com/gorilla/mux"
)

//Home handlefunc
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Home Endpoint Hit")

}

//AddPosition handlefunc
func AddPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var postBody models.PostBody
	if r.Method == "POST" {
		postBody = getJSONBodyData(w, r)
		fmt.Println(postBody.Position)
		postBody.Position.DateApplied = time.Now()
		postBody.Position.DateModified = time.Now()
		addPosition(&postBody.Position)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.Position); err != nil {
			panic(err)
		}

	}

}

//EditPosition handlefunc
func EditPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("Edit Position endpoint hit")

	if r.Method == "PUT" {
		postBody := getJSONBodyData(w, r)

		positionID, _ := strconv.Atoi(mux.Vars(r)["id"])
		postBody.Position.ID = positionID
		postBody.Position.DateModified = time.Now()

		editPosition(&postBody.Position)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.Position); err != nil {
			panic(err)
		}
	}

}

//DeletePosition handlefunc
func DeletePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		fmt.Println("Delete Position Endpoint Hit")
		positionID := mux.Vars(r)["id"]
		position := deletePosition(positionID)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(position); err != nil {
			panic(err)
		}
	}

}

//ViewPositions handlefunc
func ViewPositions(w http.ResponseWriter, r *http.Request) {
	type Position struct { //anonymous struct
		Position models.Position
		Company  models.Company
	}

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	filter := r.FormValue("filter")
	sort := r.FormValue("sort")
	if filter != "" {
		fmt.Println(filter)
	}
	if sort != "" {
		fmt.Println(sort)
	}
	filt, _ := strconv.ParseBool(filter)
	positions := viewPositions(filt, sort)
	data := []Position{}
	for _, position := range positions {
		p := Position{}

		company := viewCompany(position.CompanyID)
		p.Company = company
		p.Position = position
		data = append(data, p)
	}
	//fmt.Println(data)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

//ViewPosition handlefunc
func ViewPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	positionID, _ := strconv.Atoi(mux.Vars(r)["id"])
	position := viewPosition(positionID)
	if position.ID == 0 {
		fmt.Fprintf(w, "No Users found")
	} else {
		if err := json.NewEncoder(w).Encode(position); err != nil {
			panic(err)
		}
	}
}

/**Companies**/

//AddCompany func
func AddCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "POST" {
		postBody := getJSONBodyData(w, r)
		addCompany(&postBody.Company)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.Company); err != nil {
			panic(err)
		}
	}
}

//DeleteCompany function
func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		fmt.Println("Delete User Endpoint Hit")
		companyID := mux.Vars(r)["id"]
		company := deleteCompany(companyID)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(company); err != nil {
			panic(err)
		}
	}
}

//ViewCompany handlefunc
func ViewCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	companyID, _ := strconv.Atoi(mux.Vars(r)["id"])
	company := viewCompany(companyID)
	if company.ID == 0 {
		fmt.Fprintf(w, "No Users found")
	} else {
		if err := json.NewEncoder(w).Encode(company); err != nil {
			panic(err)
		}
	}

}

//ViewCompanies handlefunc
func ViewCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	companies := viewCompanies()

	if err := json.NewEncoder(w).Encode(companies); err != nil {
		panic(err)
	}

}

//EditCompany handlefunc
func EditCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("Edit Company endpoint hit")

	if r.Method == "PUT" {
		post := getJSONBodyData(w, r)
		fmt.Println(post)
		companyID, _ := strconv.Atoi(mux.Vars(r)["id"])
		post.Company.ID = companyID
		editCompany(&post.Company)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(post.Company); err != nil {
			panic(err)
		}
	}

}
