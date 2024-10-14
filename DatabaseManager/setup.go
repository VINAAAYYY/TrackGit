package databasemanager

import (
	"github.com/fatih/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Setup struct{}

func (s Setup) InitDb() *gorm.DB {
	color.New(color.Underline, color.FgGreen).Println("Database Initialization Started")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var migration Migration
	migration.Migrate(db)
	return db
}
