package models

import "time"

//Position Struct
type Position struct {
	//position struct
	Name         string
	Status       string
	Description  string
	ID           int
	Location     string
	Requirements string
	CompanyID    int
	DateApplied  time.Time
	DateModified time.Time
	URL          string
}
