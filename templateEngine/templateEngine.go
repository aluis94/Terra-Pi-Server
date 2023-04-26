package templateEngine

import (
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/aluis94/terra-pi-server/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var scriptsDir = "./scripts/"
var templatesDir = "./templateEngine/templates/"
var cronDir = "./cron/"

type TempJob struct {
	Job        *models.Job
	Device     *models.Device
	CondDevice *models.Device
	MsgDevice  *models.Device
}

func CreateCrontTabFile() {

}

func CreateFile(scriptName string) {
	f, err := os.Create(scriptName)
	Check(err)
	defer f.Close()

}

func DeleteFile(filename string) {
	err := os.Remove(filename)
	Check(err)
}

// Check function
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// createScript
func CreateScript(job *models.Job, device *models.Device, msgDevice *models.Device, condDevice *models.Device) string {
	tempJob := TempJob{job, device, msgDevice, condDevice}
	//generate Script name
	scriptName := CreateScriptName(&tempJob)
	//getTemplateType
	templateType := GetTemplateType(&tempJob)
	//generate script from template
	GenerateScriptFromTemplate(scriptName, templateType, &tempJob, templatesDir)

	//CreateFile(scriptName)
	return scriptName
}

func CreateScriptName(tempJob *TempJob) string {

	scriptname := scriptsDir + tempJob.Device.Name + "_" + strconv.Itoa(tempJob.Device.ID) + "_" + tempJob.Job.Name + ".py"
	//get message device and conditional device

	return scriptname
}
func GenerateScriptFromTemplate(scriptName string, templateType string, tempJob *TempJob, templDir string) {
	//create script file:output from template
	path := scriptName
	f, err := os.Create(path)
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	//parse and execute template
	templatePath := templDir + templateType
	t, err := template.ParseFiles(templatePath)
	Check(err)
	err = t.Execute(f, &tempJob)
	Check(err)
}

// get
func GetTemplateType(job *TempJob) string {

	category := ""

	extension := ".tmpl"

	//initialize caser for converting lower/title strings
	caser := cases.Title(language.Und)
	//check if there are other meeage or condition jobs
	if job.Job.CDevice_ID == 0 && job.Job.MDevice_ID == 0 {
		category = "simple"
	} else {
		category = "multi"
	}
	//Check Main Device Category: Notification, Device, Sensors
	switch mCategory := strings.ToLower(job.Device.Category); mCategory {
	case "notification":
		category = category + "Notification"
	case "sensor":
		category = category + "SensorRead"
	case "device":
		deviceOnOff := strings.ToLower(job.Job.DeviceOnOff)
		OnOff := caser.String(deviceOnOff)
		category = category + "Device" + OnOff
	default:
		category = ""
		extension = ""
		//do nothing
	}

	//return template name
	return category + extension
}

// Delete script
func DeleteScript(scriptname string) {

}

func ReadTemplateFile(path string) []byte {
	buff, err := os.ReadFile(path)
	Check(err)
	return buff
}
