package models

import (
	"time"

	"gorm.io/gorm"
)

type CommitHistory struct {
	gorm.Model
	Date       time.Time `gorm:"index:date;not null"`
	CommitHash string    `gorm:"unique;not null"`
}
