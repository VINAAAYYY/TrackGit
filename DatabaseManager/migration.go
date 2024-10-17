package databasemanager

import (
	models "TrackGit/DatabaseManager/Models"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

type Migration struct{}

func (m Migration) Migrate(db *gorm.DB) {
	color.New(color.Underline, color.FgGreen).Println("Migration For Database Started")
	err := db.AutoMigrate(&models.CommitHistory{})
	if err != nil {
		panic("Migration to Model CommitHistory Failed")
	}
}