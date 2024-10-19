package main

import (
	"log"
	"os"

	Model "TrackGit/DatabaseManager/Models"
	Repository "TrackGit/DatabaseManager/Repository"

	"github.com/fatih/color"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("No commit hash provided")
	}

	commitHash := os.Args[1]
	repoPath := os.Args[2]

	// get commit info
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("Could not open repository:", err.Error())
		return
	}

	commit, err := repo.CommitObject(plumbing.NewHash(commitHash))
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("Could not find commit:", err.Error())
		return
	}

	// connect to database
	db, err := gorm.Open(sqlite.Open("../../../commit.history.db"), &gorm.Config{})
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("Error while connecting to DB:", err.Error())
		return
	}

	// save commit
	commitObj := &Model.CommitHistory{
		Date:       commit.Author.When,
		CommitHash: commitHash,
	}
	var repository Repository.Repository
	repository.Db = db
	repository.Insert(commitObj)

}
