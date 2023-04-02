package models

//Job Stuct
type Job struct {
	ID      int `gorm:"primaryKey"`
	ChildID int
	//status
	Start_Time string
	End_Time   string
	//Script
	ScriptName  string
	Name        string
	Description string
	//Main job
	Device_ID   int
	DeviceOnOff string
	//Messaging job
	MDevice_ID int
	MCondition string
	Message    string
	//Conditional job
	Condition  string
	CDevice_ID int
	//CRON
	VerbalInstr string

	Minute    string
	Hour      string
	Day_Month string
	Month     string
	Day_Week  string
}
