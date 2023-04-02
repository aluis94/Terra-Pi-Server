package models

//Device struct
type Device struct {
	//Device struct
	ID       int `gorm:"primaryKey"`
	Name     string
	Pin1     int
	Pin2     int
	Pin3     int
	Type     string
	Category string
}
