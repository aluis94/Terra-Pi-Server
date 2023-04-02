package models

//PostBody struct
type PostBody struct {
	Devices     []Device
	Jobs        []Job
	DataEntries []DataEntry

	Device    Device
	Job       Job
	DataEntry DataEntry
}
