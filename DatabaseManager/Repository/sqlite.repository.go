package repository

import (
	model "TrackGit/DatabaseManager/Models"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (r Repository) GetAll() model.CommitHistory {
	var result model.CommitHistory
	r.Db.Find(&result)
	return result
}

func (r Repository) GetBetweenDates(startDate time.Time, endDate time.Time) model.CommitHistory {
	var result model.CommitHistory
	r.Db.Where("Date BETWEEN ? AND ?", startDate, endDate).Find(&result)
	return result
}

// Can be used for bulk insert too
func (r Repository) Insert(value interface{}) {
	r.Db.Create(value)
}

func (r Repository) Delete(hash string) {
	var commitHistory model.CommitHistory
	cond := model.CommitHistory{CommitHash: hash}
	r.Db.Where(cond).Delete(commitHistory)
}
