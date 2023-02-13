package middleware

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aluis94/terra-pi-server/models"
)

//JSON
func getJSONBodyData(w http.ResponseWriter, r *http.Request) models.PostBody {
	var post models.PostBody
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) //unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	//fmt.Println("Post: ", post)
	return post
}
