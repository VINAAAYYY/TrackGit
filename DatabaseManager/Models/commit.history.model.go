package models

import (
	"time"

	"gorm.io/gorm"
)

type CommitHistory struct {
	gorm.Model
	Date       time.Time `gorm:"index:idx_date;not null"`
	CommitHash string    `gorm:"uniqueIndex:idx_hash;not null"`
}
