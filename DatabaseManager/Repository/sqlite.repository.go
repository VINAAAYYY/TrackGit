package repository

import (
	model "TrackGit/DatabaseManager/Models"
	"sync"
	"time"

	"gorm.io/gorm"
)

// Singleton instance of Repository
var (
	instance Repository
	once     sync.Once
)

type Repository struct {
	Db *gorm.DB
}

func (r Repository) InitRepository(db *gorm.DB) {
	once.Do(func() {
		instance = Repository{Db: db}
	})
}

func (r Repository) GetRepository() Repository {
	return instance
}

func (r Repository) GetAll() ([]model.CommitHistory, error) {
	var result []model.CommitHistory
	err := r.Db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r Repository) GetBetweenDates(startDate time.Time, endDate time.Time) ([]model.CommitHistory, error) {
	var result []model.CommitHistory
	err := r.Db.Where("Date BETWEEN ? AND ?", startDate, endDate).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r Repository) Insert(value *model.CommitHistory) error {
	err := r.Db.Create(value).Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) InsertBulk(value []*model.CommitHistory) error {
	err := r.Db.Create(value).Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) Delete(hash string) {
	var commitHistory model.CommitHistory
	cond := model.CommitHistory{CommitHash: hash}
	r.Db.Where(cond).Delete(commitHistory)
}
