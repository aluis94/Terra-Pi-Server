package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJsonFile(t *testing.T) {
	fname := "temp_hum_sensor.json"
	b := readJsonFile(fname)

	t.Log(string(b))
	assert.NotEmpty(t, b)
}

func TestParseSensorDeviceFile(t *testing.T) {
	fname := "temp_hum_sensor.json"
	b := readJsonFile(fname)
	device := parseDeviceJson(b)
	t.Log(device)
	assert.NotEmpty(t, device)
}

func TestParseJobFile(t *testing.T) {
	fname := "temp_hum_job.json"
	b := readJsonFile(fname)
	device := parseJobJson(b)
	t.Log(*device)
	assert.NotEmpty(t, *device)
}

func TestLoadJsonJobs(t *testing.T) {
	jobNames := []string{"job_message_email.json", "job_message_sms.json"}
	jobs := loadJSONJobs(jobNames)
	for _, job := range *jobs {
		assert.NotEmpty(t, job)
		t.Log(job)
	}

}

func TestJsonStringFormat(t *testing.T) {

}
