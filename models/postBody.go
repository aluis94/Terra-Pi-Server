package models

//PostBody struct
type PostBody struct {
	Device Device
	Sensor Sensor
	Jobs   []Job
}
