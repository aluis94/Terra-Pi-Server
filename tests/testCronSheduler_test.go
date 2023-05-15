package tests

import (
	"testing"
	"time"

	"github.com/aluis94/terra-pi-server/cron"
	"github.com/aluis94/terra-pi-server/models"
	"github.com/stretchr/testify/assert"
)

func TestExecEmailScript(t *testing.T) {
	out := cron.ExecScript("sendEmail.py")
	t.Log(out)
	assert.NotEmpty(t, out)
}

func TestExecSMSScript(t *testing.T) {
	out := cron.ExecScript("sendSMS.py")
	t.Log(out)
	assert.NotEmpty(t, out)
}
func TestSimpleJob(t *testing.T) {
	var jobs []models.Job
	s, _ := cron.RunCronJobs(&jobs)
	time.Sleep(5 * time.Second)
	s.Stop()
	assert.True(t, true)
}

func TestCronScheduler(t *testing.T) {
	var jobs []models.Job
	b := readJsonFile("job_message_email.json")
	job := parseJobJson(b)
	jobs = append(jobs, *job)

	b = readJsonFile("job_message_sms.json")
	job = parseJobJson(b)
	jobs = append(jobs, *job)

	cron.RunCronJobs(&jobs)

}

func TestCreateCronExpression(t *testing.T) {

	b := readJsonFile("job_message_email.json")
	job := parseJobJson(b)

	expr := cron.CreateCronExpression(job)
	assert.Equal(t, "*/1 * * * *", expr)
}
