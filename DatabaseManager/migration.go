package databasemanager

import (
	models "GitHistoryTracker/DatabaseManager/Models"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

type Migration struct{}

func (m Migration) Migrate(db *gorm.DB) {
	color.New(color.Underline, color.FgGreen).Println("Migration For Database Started")
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Migration to Model User Failed")
	}
	err1 := db.AutoMigrate(&models.CommitHistory{})
	if err1 != nil {
		panic("Migration to Model CommitHistory Failed")
	}
}
