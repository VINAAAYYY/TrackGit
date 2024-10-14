package main

import (
	DbManager "GitHistoryTracker/DatabaseManager"
	Repository "GitHistoryTracker/DatabaseManager/Repository"
	TrackGirDirs "GitHistoryTracker/GitDirectorySearch"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.New(color.Italic, color.FgGreen).Println("Initializing...")
	var Setup DbManager.Setup
	db := Setup.InitDb()

	var trackRetroCommits TrackGirDirs.RetroCommitTracker
	commitHistory := trackRetroCommits.Track(db)

	var repository Repository.Repository
	repository.Db = db
	repository.Insert(commitHistory)

	// get commit history (test only)
	dbCommitHistory := repository.GetAll()
	fmt.Println(dbCommitHistory)
}
