package templateEngine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/aluis94/terra-pi-server/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
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

func DeleteFile(dir string, filename string) {
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

func ReplaceSpaceWithUnderscore(name string) string {
	formatted := ""
	formatted = strings.Replace(name, " ", "_", -1)

	return formatted
}

func CreateScriptName(tempJob *TempJob) string {
	jobName := ReplaceSpaceWithUnderscore(tempJob.Job.Name)
	scriptname := scriptsDir + tempJob.Device.Name + "_" + strconv.Itoa(tempJob.Device.ID) + "_" + jobName + ".py"
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
	defer f.Close()
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
		return category + extension
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

func WriteData() {
	// HTTP endpoint
	fname := "./scripts/data.txt"
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect DataEntrybase")
	}

	defer db.Close()

	readFile, err := os.Open(fname)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	// Create a HTTP post request
	var de models.DataEntry
	db.Last(&de)
	fmt.Println(de.ID)
	currID := de.ID + 1
	var Entry *models.PostBody

	for fileScanner.Scan() {
		dataEntry := models.DataEntry{}
		fmt.Println(strings.Replace(fileScanner.Text(), "'", "\"", -1))
		err = json.Unmarshal([]byte(strings.Replace(fileScanner.Text(), "'", "\"", -1)), &Entry)
		if err != nil {
			fmt.Println("Error during Unmarshal(): ", err)
		}
		dataEntry.ID = currID
		dataEntry.Device_ID = Entry.DataEntry.Device_ID
		dataEntry.Type = Entry.DataEntry.Type
		dataEntry.Value = Entry.DataEntry.Value
		dataEntry.Unit = Entry.DataEntry.Unit
		currTime := time.Now().Format("2006-01-02 15:04:05")
		dataEntry.TimeStamp = currTime
		fmt.Println("Entry before:", dataEntry)
		db.Create(dataEntry)
		fmt.Println("DataEntry created")

		fmt.Println("Entry:", dataEntry)
		currID = currID + 1

	}
	defer CreateFile(fname)
	defer readFile.Close()

}
