package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aluis94/terra-pi-server/models"
	"github.com/gorilla/mux"
)

// Home handlefunc
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Home Endpoint Hit")

}

/**Devices**/
// AddDevice handlefunc
func AddDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var postBody models.PostBody
	if r.Method == "POST" {
		postBody = getJSONBodyData(w, r)
		fmt.Println(postBody.Device)
		createDevice(&postBody.Device)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.Device); err != nil {
			panic(err)
		}

	}

}

// EditDevice handlefunc
func EditDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("Edit Device endpoint hit")

	if r.Method == "PUT" {
		postBody := getJSONBodyData(w, r)

		deviceID, _ := strconv.Atoi(mux.Vars(r)["id"])
		postBody.Device.ID = deviceID
		fmt.Println(postBody)
		editDevice(&postBody.Device)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.Device); err != nil {
			panic(err)
		}
	}

}

// DeleteDevice handlefunc
func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		fmt.Println("Delete Device Endpoint Hit")
		deviceID := mux.Vars(r)["id"]
		device := deleteDevice(deviceID)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(device); err != nil {
			panic(err)
		}
	}

}

// ViewDevice handlefunc
func ViewDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	deviceID, _ := strconv.Atoi(mux.Vars(r)["id"])
	device := GetDevice(deviceID)
	if device.ID == 0 {
		fmt.Fprintf(w, "No Devices found")
	} else {
		if err := json.NewEncoder(w).Encode(device); err != nil {
			panic(err)
		}
	}
}

// ViewDevices handlefunc
func ViewDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	category := strings.ToLower(r.FormValue("category"))
	devices := viewDevices(category)

	if err := json.NewEncoder(w).Encode(devices); err != nil {
		panic(err)
	}

}

/**Data Entries**/
// Add data entry
func AddDataEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET, OPTIONS,")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var postBody models.PostBody
	if r.Method == "POST" {
		postBody = getJSONBodyData(w, r)
		fmt.Println(postBody.DataEntry)
		createDataEntry(&postBody.DataEntry)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.DataEntry); err != nil {
			panic(err)
		}

	}

}

// EditDataEntry handlefunc
func EditDataEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("Edit DataEntry endpoint hit")

	if r.Method == "PUT" {
		postBody := getJSONBodyData(w, r)

		DataEntryID, _ := strconv.Atoi(mux.Vars(r)["id"])
		postBody.DataEntry.ID = DataEntryID
		fmt.Println(postBody)
		editDataEntry(&postBody.DataEntry)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.DataEntry); err != nil {
			panic(err)
		}
	}

}

// DeleteDataEntry handlefunc
func DeleteDataEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		fmt.Println("Delete DataEntry Endpoint Hit")
		DataEntryID := mux.Vars(r)["id"]
		DataEntry := deleteDataEntry(DataEntryID)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(DataEntry); err != nil {
			panic(err)
		}
	}

}

// ViewDataEntry handlefunc
func ViewDataEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	DataEntryID, _ := strconv.Atoi(mux.Vars(r)["id"])
	DataEntry := viewDataEntry(DataEntryID)
	if DataEntry.ID == 0 {
		fmt.Fprintf(w, "No DataEntrys found")
	} else {
		if err := json.NewEncoder(w).Encode(DataEntry); err != nil {
			panic(err)
		}
	}
}

// ViewDataEntries handlefunc
func ViewDataEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	DataEntrys := viewDataEntries()

	if err := json.NewEncoder(w).Encode(DataEntrys); err != nil {
		panic(err)
	}

}

/**Jobs**/
// Add job
func AddJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var postBody models.PostBody
	if r.Method == "POST" {
		postBody = getJSONBodyData(w, r)
		fmt.Println(postBody.Job)
		createJob(&postBody.Job)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.Job); err != nil {
			panic(err)
		}

	}

}

// EditJob handlefunc
func EditJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("Edit Job endpoint hit")

	if r.Method == "PUT" {
		postBody := getJSONBodyData(w, r)

		JobID, _ := strconv.Atoi(mux.Vars(r)["id"])
		postBody.Job.ID = JobID
		fmt.Println(postBody)
		editJob(&postBody.Job)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(postBody.Job); err != nil {
			panic(err)
		}
	}

}

// DeleteJob handlefunc
func DeleteJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		fmt.Println("Delete Job Endpoint Hit")
		JobID := mux.Vars(r)["id"]
		Job := deleteJob(JobID)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(Job); err != nil {
			panic(err)
		}
	}

}

// ViewJob handlefunc
func ViewJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "GET" {
		JobID, _ := strconv.Atoi(mux.Vars(r)["id"])
		Job := viewJob(JobID)
		if Job.ID == 0 {
			fmt.Fprintf(w, "No Jobs found")
		} else {
			if err := json.NewEncoder(w).Encode(Job); err != nil {
				panic(err)
			}
		}
	}
}

// ViewJobs handlefunc
func ViewJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	Jobs := viewJobs()
	if r.Method == "GET" {
		if err := json.NewEncoder(w).Encode(Jobs); err != nil {
			panic(err)
		}
	}

	if r.Method == "OPTIONS" {
		if err := json.NewEncoder(w).Encode(Jobs); err != nil {
			panic(err)
		}
	}

}
