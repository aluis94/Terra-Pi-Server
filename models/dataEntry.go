package models

// Data Entry struct
type DataEntry struct {
	//DataEntry struct
	ID        int `gorm:"primaryKey"`
	Device_ID int
	Type      string
	TimeStamp string
	Value     float32
	Unit      string
}
