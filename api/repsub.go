package repsub

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "fmt"

)

type Report struct {
	ID int
	Name string
	Description string
	Severity string
	HackerID int
	Hacker Hacker 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Hacker struct {
	ID uint
	Username string
	Rank string
}

func ConnectDB() (db *gorm.DB) {
	dsn := "homestead:secret@tcp(192.168.5.10:3306)/db_repsub?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
    	panic("failed to connect database")
  	}
	  
	// db.AutoMigrate(&Hacker{},&Report{})

	return db
}

func SubmitReport(name string,desc string, severity string, hacker string) string {
	message := fmt.Sprintf("Thank you for your report! Your report has been validated and you've been awarded %s of bounty!Congratulations %s!", "500", hacker)
	
	database := ConnectDB()

	report := Report{Name: name, Description: desc, Severity: severity, HackerID: 1}

	database.Create(&report)
	
	return message
}

func ViewReports() {

	database := ConnectDB()
	reportObj, err := database.Where("hacker_id = ?", "1").Find(&reports)
	
	report := Reports{}
    results := []Reports{}

    for reportObj.Next() {
        var id int
        var name, description, severity string
        err = reportObj.Scan(&name, &description, &severity)

        if err != nil {
            panic(err.Error())
        }
        report.Name = name
        report.Description = description
        report.Severity = severity
        results = append(results, report)
    }

	return result
}


