package postcheckout

import (
	"log"
	"os"

	Repository "TrackGit/DatabaseManager/Repository"

	"github.com/fatih/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No commit hashes provided")
	}

	commitHashes := os.Args[1:]

	// connect to database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("Error while connecting to DB:", err.Error())
		return
	}

	var repository Repository.Repository
	repository.Db = db

	// delete commits
	for _, commitHash := range commitHashes {
		repository.Delete(commitHash)
	}
}
