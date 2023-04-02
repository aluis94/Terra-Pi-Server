package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/aluis94/terra-pi-server/models"
	"github.com/aluis94/terra-pi-server/templateEngine"
	"github.com/stretchr/testify/assert"
)

var testScriptsDir = "./scripts/"
var testTemplatesDir = "../templateEngine/templates/"
var currTemplateDir = "./scripts/"

func TestCreateFile(t *testing.T) {
	scriptname := "test.txt"

	templateEngine.CreateFile(testScriptsDir + scriptname)

	assert.FileExists(t, testScriptsDir+scriptname, "file not found %s")
}

func TestDeleteFile(t *testing.T) {
	scriptname := "test.txt"

	templateEngine.DeleteFile(testScriptsDir + scriptname)

	assert.NoFileExists(t, testScriptsDir+scriptname, "file still exists %s")
}

func TestGetWorkingDirectory(t *testing.T) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	t.Log("This is curr dir:\n " + mydir)
	assert.True(t, true)
}

func TestReadTemplate(t *testing.T) {

	buff := templateEngine.ReadTemplateFile(testTemplatesDir + "test.tmpl")
	t.Log("This is curr dir:\n " + string(buff))
	assert.NotEmpty(t, buff)
}

func TestGenerateTemplate(t *testing.T) {
	job := models.Job{}
	job.ID = 1
	job.Name = "test"
	job.DeviceOnOff = "ON"
	device := models.Device{}
	device.ID = 1
	device.Name = "device"
	device.Pin1 = 1
	device.Pin2 = 2
	device.Pin3 = 0
	device.Category = "Device"
	device.Type = "Heater"
	cDevice := models.Device{}
	mDevice := models.Device{}
	tempJob := templateEngine.TempJob{Job: &job, Device: &device, CondDevice: &mDevice, MsgDevice: &cDevice}
	scriptName := templateEngine.CreateScriptName(&tempJob)
	templateType := templateEngine.GetTemplateType(&tempJob)
	templateEngine.GenerateScriptFromTemplate(scriptName, templateType, &tempJob, testTemplatesDir)
	assert.Equal(t, currTemplateDir+"device_1_test.py", scriptName)
}

func TestSimpleTemp_HumTemplate(t *testing.T) {
	job := models.Job{}
	job.ID = 1
	job.Name = "simple_temp_sensor_test"
	job.DeviceOnOff = "ON"
	device := models.Device{}
	device.ID = 1
	device.Name = "DHT11"
	device.Pin1 = 21
	device.Pin2 = 2
	device.Pin3 = 0
	device.Category = "Sensor"
	device.Type = "Temp_Hum"
	cDevice := models.Device{}
	mDevice := models.Device{}
	tempJob := templateEngine.TempJob{Job: &job, Device: &device, CondDevice: &mDevice, MsgDevice: &cDevice}
	scriptName := templateEngine.CreateScriptName(&tempJob)
	templateType := templateEngine.GetTemplateType(&tempJob)
	templateEngine.GenerateScriptFromTemplate(scriptName, templateType, &tempJob, testTemplatesDir)
	assert.Equal(t, currTemplateDir+"DHT11_1_simple_temp_sensor_test.py", scriptName)
}

func TestSimpleMotionTemplate(t *testing.T) {
	job := models.Job{}
	job.ID = 1
	job.Name = "simple_motion_sensor_test"
	job.DeviceOnOff = "ON"
	device := models.Device{}
	device.ID = 1
	device.Name = "UltraSonic"
	device.Pin1 = 26
	device.Pin2 = 21
	device.Pin3 = 0
	device.Category = "Sensor"
	device.Type = "Motion"
	cDevice := models.Device{}
	mDevice := models.Device{}
	tempJob := templateEngine.TempJob{Job: &job, Device: &device, CondDevice: &mDevice, MsgDevice: &cDevice}
	scriptName := templateEngine.CreateScriptName(&tempJob)
	templateType := templateEngine.GetTemplateType(&tempJob)
	templateEngine.GenerateScriptFromTemplate(scriptName, templateType, &tempJob, testTemplatesDir)
	assert.Equal(t, currTemplateDir+"UltraSonic_1_simple_motion_sensor_test.py", scriptName)
}

func TestSimpleLDRTemplate(t *testing.T) {
	job := models.Job{}
	job.ID = 1
	job.Name = "simple_light_sensor_test"
	job.DeviceOnOff = "ON"
	device := models.Device{}
	device.ID = 1
	device.Name = "LDR"
	device.Pin1 = 26
	device.Pin2 = 21
	device.Pin3 = 0
	device.Category = "Sensor"
	device.Type = "Light"
	cDevice := models.Device{}
	mDevice := models.Device{}
	tempJob := templateEngine.TempJob{Job: &job, Device: &device, CondDevice: &mDevice, MsgDevice: &cDevice}
	scriptName := templateEngine.CreateScriptName(&tempJob)
	templateType := templateEngine.GetTemplateType(&tempJob)
	templateEngine.GenerateScriptFromTemplate(scriptName, templateType, &tempJob, testTemplatesDir)
	assert.Equal(t, currTemplateDir+"LDR_1_simple_light_sensor_test.py", scriptName)
}

func TestSimpleEmailTemplate(t *testing.T) {
	dfname := "device_email.json"
	b1 := readJsonFile(dfname)
	device := parseDeviceJson(b1)

	jfname := "job_message_email.json"
	b2 := readJsonFile(jfname)
	job := parseJobJson(b2)

	cDevice := models.Device{}
	mDevice := models.Device{}
	tempJob := templateEngine.TempJob{Job: job, Device: device, CondDevice: &mDevice, MsgDevice: &cDevice}
	scriptName := templateEngine.CreateScriptName(&tempJob)
	templateType := templateEngine.GetTemplateType(&tempJob)
	templateEngine.GenerateScriptFromTemplate(scriptName, templateType, &tempJob, testTemplatesDir)
	assert.Equal(t, currTemplateDir+"Email_1_TestEmail.py", scriptName)
}

func TestSimpleSMSTemplate(t *testing.T) {
	dfname := "device_sms.json"
	b1 := readJsonFile(dfname)
	device := parseDeviceJson(b1)

	jfname := "job_message_sms.json"
	b2 := readJsonFile(jfname)
	job := parseJobJson(b2)

	cDevice := models.Device{}
	mDevice := models.Device{}
	tempJob := templateEngine.TempJob{Job: job, Device: device, CondDevice: &mDevice, MsgDevice: &cDevice}
	scriptName := templateEngine.CreateScriptName(&tempJob)
	templateType := templateEngine.GetTemplateType(&tempJob)
	templateEngine.GenerateScriptFromTemplate(scriptName, templateType, &tempJob, testTemplatesDir)
	assert.Equal(t, currTemplateDir+"Text_2_TestSMS.py", scriptName)
}

func TestSimpleDeviceOffTemplate(t *testing.T) {
	dfname := "device_light.json"
	b1 := readJsonFile(dfname)
	device := parseDeviceJson(b1)

	jfname := "job_device_off.json"
	b2 := readJsonFile(jfname)
	job := parseJobJson(b2)

	cDevice := models.Device{}
	mDevice := models.Device{}
	tempJob := templateEngine.TempJob{Job: job, Device: device, CondDevice: &mDevice, MsgDevice: &cDevice}
	scriptName := templateEngine.CreateScriptName(&tempJob)
	templateType := templateEngine.GetTemplateType(&tempJob)
	templateEngine.GenerateScriptFromTemplate(scriptName, templateType, &tempJob, testTemplatesDir)
	assert.Equal(t, currTemplateDir+"Light1_8_TestDeviceOff.py", scriptName)
}

func TestSimpleDeviceONTemplate(t *testing.T) {
	dfname := "device_light.json"
	b1 := readJsonFile(dfname)
	device := parseDeviceJson(b1)

	jfname := "job_device_on.json"
	b2 := readJsonFile(jfname)
	job := parseJobJson(b2)

	cDevice := models.Device{}
	mDevice := models.Device{}
	tempJob := templateEngine.TempJob{Job: job, Device: device, CondDevice: &mDevice, MsgDevice: &cDevice}
	scriptName := templateEngine.CreateScriptName(&tempJob)
	templateType := templateEngine.GetTemplateType(&tempJob)
	templateEngine.GenerateScriptFromTemplate(scriptName, templateType, &tempJob, testTemplatesDir)
	assert.Equal(t, currTemplateDir+"Light1_8_TestDeviceOn.py", scriptName)
}
