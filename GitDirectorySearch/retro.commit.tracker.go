package gitdirectorysearch

import (
	color "github.com/fatih/color"

	model "GitHistoryTracker/DatabaseManager/Models"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"gorm.io/gorm"
)

type RetroCommitTracker struct{}

func (rct RetroCommitTracker) Track(*gorm.DB) []model.CommitHistory {
	var search Search
	history := []model.CommitHistory{}
	gitDirectories := search.TrackGirDirs()
	for _, dir := range gitDirectories {
		color.New(color.FgBlue).Println("Checking Commits in", dir)
		r, err := git.PlainOpen(dir)
		if err != nil {
			color.New(color.Bold, color.FgRed).Println("Failed to open git repo:", err.Error())
			continue
		}

		ref, err := r.Head()
		if err != nil {
			color.New(color.Bold, color.FgRed).Println("Failed to get repo log:", err.Error())
			continue
		}

		commitIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
		if err != nil {
			color.New(color.Bold, color.FgRed).Println("Error while fetching commit log:", err.Error())
			continue
		}

		err = commitIter.ForEach(func(c *object.Commit) error {
			// replace with user email
			if c.Committer.Email == "vinay.singhal43@gmail.com" {
				history = append(history, model.CommitHistory{
					Date:       c.Author.When,
					CommitHash: c.Hash.String(),
				})
			}
			return nil
		})

		if err != nil {
			color.New(color.Bold, color.FgRed).Println("Error while iterating commits:", err.Error())
		}

	}
	return history
}
