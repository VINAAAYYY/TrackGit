package main

import (
	DbManager "TrackGit/DatabaseManager"
	Repository "TrackGit/DatabaseManager/Repository"
	TrackGirDirs "TrackGit/GitDirectorySearch"

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

	color.New(color.Italic, color.FgGreen).Println("Initialization Completed 🥳")
}
