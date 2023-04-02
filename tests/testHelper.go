package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aluis94/terra-pi-server/models"
)

var mockDataPath = "./mockdata/"

// check err function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readJsonFile(fname string) []byte {
	// Open our jsonFile
	jsonFile, err := os.Open(mockDataPath + fname)
	// if we os.Open returns an error then handle it
	check(err)
	fmt.Println("Successfully Opened " + fname + " file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	check(err)
	return b
}

func parseDeviceJson(b []byte) *models.Device {
	var postBody models.PostBody
	if err := json.Unmarshal(b, &postBody); err != nil {
		fmt.Println("Error with unmarshal")
	}

	fmt.Println(postBody.Device.Name)
	return &postBody.Device
}

func parseJobJson(b []byte) *models.Job {
	var postBody models.PostBody
	if err := json.Unmarshal(b, &postBody); err != nil {
		fmt.Println("Error with unmarshal")
	}

	fmt.Println(postBody.Device.Name)
	return &postBody.Job
}

func loadJSONJobs(jobNames []string) *[]models.Job {
	var jobs []models.Job
	for _, jobName := range jobNames {
		b := readJsonFile(jobName)
		job := parseJobJson(b)
		jobs = append(jobs, *job)
	}
	return &jobs
}
