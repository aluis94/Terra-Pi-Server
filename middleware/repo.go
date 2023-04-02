package middleware

import (
	"fmt"
	"time"

	"github.com/aluis94/terra-pi-server/cron"
	"github.com/aluis94/terra-pi-server/models"
	"github.com/aluis94/terra-pi-server/templateEngine"
	"github.com/go-co-op/gocron"

	"github.com/jinzhu/gorm"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var scheduler *gocron.Scheduler

// InitialMigration gorm
func InitialMigration() {
	//scheduler

	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Job{}, &models.DataEntry{}, &models.Device{})

	//create the email, text, and push notification default devices
	email := models.Device{Name: "Email", Type: "Email", Category: "Notification"}
	text := models.Device{Name: "Text", Type: "SMS", Category: "Notification"}
	push := models.Device{Name: "Push", Type: "Push", Category: "Notification"}
	var Devices []models.Device
	query := "select * from Devices where id in (1,2,3)"
	rows, err := db.Raw(query).Rows() // (*sql.Rows, error)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var Device models.Device
		// ScanRows scan a row into user
		db.ScanRows(rows, &Device)
		Devices = append(Devices, Device)
	}
	if len(Devices) == 0 {

		db.Create(&email)
		db.Create(&text)
		db.Create(&push)
	}

	//startScheduler
	//get all jobs
	jobs := viewJobs()
	if len(jobs) != 0 { //if there is at least one job
		fmt.Println("Starting scheduler")
		scheduler = cron.RunCronJobs(&jobs)
	}
}

/**Devices**/

// CreateDevice
func createDevice(Device *models.Device) {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()

	db.Create(Device)
	fmt.Println("Device created")
}

// Edit Device
func editDevice(Device *models.Device) {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()

	var dbDevice models.Device
	db.Where("id = ?", Device.ID).Find(&dbDevice)

	//update Device DataEntry
	dbDevice.Name = Device.Name
	dbDevice.Pin1 = Device.Pin1
	dbDevice.Pin2 = Device.Pin2
	dbDevice.Pin3 = Device.Pin3
	dbDevice.Type = Device.Type
	dbDevice.Category = Device.Category

	if dbDevice.ID != 0 && dbDevice.Name != "" && dbDevice.Type != "" && dbDevice.Category != "" {
		db.Save(&dbDevice)
		fmt.Println("Successfully Updated Device")
	} else {
		fmt.Println("Device does not exist")
	}

}

//Delete Device

func deleteDevice(id string) models.Device {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()

	var device models.Device
	db.Where("id = ?", id).Find(&device)
	if device.ID != 0 {
		db.Delete(&device)
		fmt.Println("Successfully Deleted Device")
	} else {
		fmt.Println("No Device deleted")
	}

	return device
}

// view Devices
func viewDevices() []models.Device {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var devices []models.Device
	db.Find(&devices)
	//fmt.Println("{}", Devices)
	return devices
}

// view single Device by ID
func GetDevice(id int) models.Device {
	var Device models.Device
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()
	db.Where("id = ?", id).Find(&Device)
	return Device
}

/**DataEntries*/

// Create Data Entry
func createDataEntry(DataEntry *models.DataEntry) {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()
	myDataEntry := DataEntry
	currTime := time.Now().Format("2006-01-02 15:04:05")
	myDataEntry.TimeStamp = currTime
	db.Create(DataEntry)
	fmt.Println("DataEntry created")
}

// Delete DataEntry
func deleteDataEntry(id string) models.DataEntry {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()

	var DataEntry models.DataEntry
	db.Where("id = ?", id).Find(&DataEntry)
	if DataEntry.ID != 0 {
		db.Delete(&DataEntry)
		fmt.Println("Successfully Deleted DataEntry")
	} else {
		fmt.Println("No DataEntry Deleted")
	}

	return DataEntry
}

// view single DataEntry by ID
func viewDataEntry(id int) models.DataEntry {
	var DataEntry models.DataEntry
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()
	db.Where("id = ?", id).Find(&DataEntry)
	return DataEntry
}

// viewDataEntries
func viewDataEntries() []models.DataEntry {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()
	var DataEntries []models.DataEntry
	db.Find(&DataEntries)
	//fmt.Println("{}", DataEntries)
	return DataEntries
}

// editDataEntry
func editDataEntry(DataEntry *models.DataEntry) {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()

	var dbDataEntry models.DataEntry
	db.Where("id = ?", DataEntry.ID).Find(&dbDataEntry)

	//update DataEntry DataEntry
	currTime := time.Now().Format("2006-01-02 15:04:05")

	dbDataEntry.Device_ID = DataEntry.Device_ID
	dbDataEntry.Value = DataEntry.Value
	dbDataEntry.Unit = DataEntry.Unit
	dbDataEntry.TimeStamp = currTime

	if dbDataEntry.ID != 0 {
		db.Save(&dbDataEntry)
		fmt.Println("Successfully Updated DataEntry")
	} else {
		fmt.Println("DataEntry does not exist")
	}

}

/**Jobs*/

// Create Job
func createJob(Job *models.Job) {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()
	//check if Device exists
	var myDevice models.Device
	fmt.Println(Job.Device_ID)
	fmt.Println(myDevice)
	db.Where("id = ?", Job.Device_ID).Find(&myDevice)
	if myDevice.ID != 0 {
		//createscript
		msgDevice := models.Device{}
		condDevice := models.Device{}
		//check if other devices exist
		if Job.MDevice_ID != 0 {
			db.Where("id = ?", Job.Device_ID).Find(&msgDevice)
		}
		if Job.CDevice_ID != 0 {
			db.Where("id = ?", Job.Device_ID).Find(&condDevice)
		}

		scriptName := templateEngine.CreateScript(Job, &myDevice, &msgDevice, &condDevice)
		Job.ScriptName = scriptName
		db.Create(Job)
		fmt.Println("Job created")

		//create and schedule script

	} else {
		fmt.Println("No exisiting device, no job created")
	}

}

// delete Job
func deleteJob(id string) models.Job {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()

	var Job models.Job
	db.Where("id = ?", id).Find(&Job)
	db.Unscoped().Delete(&Job)

	fmt.Println("Successfully Deleted Job")
	return Job
}

// view single Job by ID
func viewJob(id int) models.Job {
	var Job models.Job
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()
	db.Where("id = ?", id).Find(&Job)
	return Job
}

// viewJobs
func viewJobs() []models.Job {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()
	var Jobs []models.Job
	db.Find(&Jobs)
	//fmt.Println("{}", Jobs)
	return Jobs
}

// editJob
func editJob(Job *models.Job) {
	db, err := gorm.Open("sqlite3", "terra-pi.db")
	if err != nil {
		panic("failed to connect DataEntrybase")
	}
	defer db.Close()

	var dbJob models.Job
	db.Where("id = ?", Job.ID).Find(&dbJob)

	//update Job DataEntry
	//status
	dbJob.Start_Time = Job.Start_Time
	dbJob.End_Time = Job.End_Time
	//script
	dbJob.Name = Job.Name
	dbJob.Description = Job.Description
	dbJob.ScriptName = Job.ScriptName
	//Main job
	dbJob.Device_ID = Job.Device_ID
	//Messaging job
	dbJob.MDevice_ID = Job.MDevice_ID
	dbJob.MCondition = Job.MCondition
	dbJob.Message = Job.Message
	//Conditional job
	dbJob.CDevice_ID = Job.CDevice_ID
	dbJob.Condition = Job.Condition
	//CRON
	dbJob.Minute = Job.Minute
	dbJob.Hour = Job.Minute
	dbJob.Day_Month = Job.Minute
	dbJob.Month = Job.Minute
	dbJob.Day_Week = Job.Minute

	if dbJob.ID != 0 {
		db.Save(&dbJob)
		fmt.Println("Successfully Updated Job")
	} else {
		fmt.Println("Job does not exist")
	}

}
