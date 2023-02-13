package middleware

import (
	"fmt"

	"github.com/aluis94/terra-pi-server/models"
	"github.com/jinzhu/gorm"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

//InitialMigration gorm
func InitialMigration() {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Position{}, &models.Company{})
}

/**Positions*/

//Add Position
func addPosition(position *models.Position) {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(position)
	fmt.Println("position created")
}

//editposition
func editPosition(position *models.Position) {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var dbPosition models.Position
	db.Where("id = ?", position.ID).Find(&dbPosition)

	//update position data
	dbPosition.Name = position.Name
	dbPosition.Status = position.Status
	dbPosition.Description = position.Description
	dbPosition.Location = position.Location
	dbPosition.Requirements = position.Requirements
	fmt.Println("CompanyID: ", position.CompanyID)
	dbPosition.CompanyID = position.CompanyID

	dbPosition.URL = position.URL
	dbPosition.DateModified = position.DateModified

	if dbPosition.ID != 0 {
		db.Save(&dbPosition)
		fmt.Println("Successfully Updated position")
	} else {
		fmt.Println("position does not exist")
	}

}

//delete position

func deletePosition(id string) models.Position {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var position models.Position
	db.Where("id = ?", id).Find(&position)
	db.Unscoped().Delete(&position)

	fmt.Println("Successfully Deleted position")
	return position
}

//view al Positions

func viewPositions(filter bool, sort string) []models.Position {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var positions []models.Position
	query := "select * from positions"
	if filter || sort != "" {
		if filter && sort != "" {
			query = query + " where status not in ('NA','Rejected','No Callback') order by " + sort
		} else if !filter && sort != "" {
			query = query + " order by " + sort
		} else if filter && sort == "" {
			query = query + " where status not in ('NA','Rejected','No Callback')"
		}
		fmt.Println(query)
		rows, err := db.Raw(query).Rows() // (*sql.Rows, error)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			var position models.Position
			// ScanRows scan a row into user
			db.ScanRows(rows, &position)
			positions = append(positions, position)
		}
	} else {
		db.Find(&positions)
	}

	//fmt.Println("{}", Positions)
	return positions
}

//view single Position by ID
func viewPosition(id int) models.Position {
	var position models.Position
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.Where("id = ?", id).Find(&position)
	return position
}

/**Companies*/

//Add Company
func addCompany(company *models.Company) {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(company)
	fmt.Println("company created")
}

//delete Company

func deleteCompany(id string) models.Company {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var company models.Company
	db.Where("id = ?", id).Find(&company)
	db.Unscoped().Delete(&company)

	fmt.Println("Successfully Deleted Company")
	return company
}

//view single Company by ID
func viewCompany(id int) models.Company {
	var company models.Company
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.Where("id = ?", id).Find(&company)
	return company
}

//view Companies
func viewCompanies() []models.Company {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var companies []models.Company
	db.Find(&companies)
	//fmt.Println("{}", companies)
	return companies
}

//editCompany
func editCompany(company *models.Company) {
	db, err := gorm.Open("sqlite3", "job.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var dbCompany models.Company
	db.Where("id = ?", company.ID).Find(&dbCompany)

	//update Company data
	dbCompany.Name = company.Name
	dbCompany.About = company.About

	if dbCompany.ID != 0 {
		db.Save(&dbCompany)
		fmt.Println("Successfully Updated Company")
	} else {
		fmt.Println("Company does not exist")
	}

}
