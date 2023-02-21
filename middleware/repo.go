package middleware

import (
	"fmt"

	"github.com/aluis94/terra-pi-server/models"
	"github.com/jinzhu/gorm"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// InitialMigration gorm
func InitialMigration() {
	db, err := gorm.Open("sqlite3", "jobs.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Device{}, &models.Sensor{}, &models.Job{})
}
