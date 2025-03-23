package contributionvisualizer

import (
	models "TrackGit/DatabaseManager/Models"
	Repository "TrackGit/DatabaseManager/Repository"
	"time"

	"github.com/fatih/color"
	"github.com/rivo/tview"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ContributionVisualizer struct {
	app     *tview.Application
	tooltip *tview.Box
	table   *tview.Table
}

func (c ContributionVisualizer) GetCommitCount(startDate time.Time) []models.CommitHistory {
	db, err := gorm.Open(sqlite.Open("../../commit.history.db"), &gorm.Config{})
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("Error while connecting to DB:", err.Error())
		return nil
	}
	var repo Repository.Repository
	repo.Db = db
	commits, err := repo.GetBetweenDates(startDate, time.Now())
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("Error while fetching commit count", err.Error())
	}
	return commits
}
