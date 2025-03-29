package gitdirectorysearch

import (
	"context"
	"slices"
	"strings"

	color "github.com/fatih/color"

	model "TrackGit/DatabaseManager/Models"

	exec "os/exec"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/sethvargo/go-envconfig"
	"gorm.io/gorm"
)

type RetroCommitTracker struct{}

type Config struct {
	OTHER_USER_PROFILES []string `env:"OTHER_USER_PROFILES"`
}

func (rct RetroCommitTracker) Track(*gorm.DB) []*model.CommitHistory {
	color.New(color.Underline, color.FgGreen).Println("Commit Tracking Started")
	var search Search
	history := []*model.CommitHistory{}
	gitDirectories := search.TrackGitDirs()
	ctx := context.Background()

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		color.New(color.Bold, color.FgYellow).Println("Failed load environment file:", err.Error())
	}
	emailsToTrack := c.OTHER_USER_PROFILES
	emailsToTrack = append(emailsToTrack, rct.getGitGlobalConfigEmail())

	for _, dir := range gitDirectories {
		color.New(color.FgBlue).Println("Checking Commits in", dir)
		repo, err := git.PlainOpen(dir)
		if err != nil {
			color.New(color.Bold, color.FgRed).Println("Failed to open git repo:", err.Error())
			continue
		}

		ref, err := repo.Head()
		if err != nil {
			color.New(color.Bold, color.FgRed).Println("Failed to get repo log:", err.Error())
			continue
		}

		commitIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
		if err != nil {
			color.New(color.Bold, color.FgRed).Println("Error while fetching commit log:", err.Error())
			continue
		}

		repoUserEmail := rct.getGitRepoConfigEmail(repo)
		emailsToTrack = append(emailsToTrack, repoUserEmail)
		err = commitIter.ForEach(func(c *object.Commit) error {
			if slices.Contains(emailsToTrack, c.Committer.Email) {
				history = append(history, &model.CommitHistory{
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
	color.New(color.Underline, color.FgGreen).Println("Commit Tracking Completed")
	return history
}

func (RetroCommitTracker) getGitRepoConfigEmail(repo *git.Repository) string {
	repoConfig, err := repo.Config()
	if err != nil {
		color.New(color.Bold, color.FgYellow).Println(
			"Failed to fetch repository config. Only global and additionally provided emails will be tracked:",
			err.Error())
	}

	repoUserEmail := repoConfig.User.Email
	return repoUserEmail
}

func (RetroCommitTracker) getGitGlobalConfigEmail() string {
	cmd := exec.Command("git", "config", "--global", "user.email")
	output, err := cmd.Output()
	if err != nil || string(output) == "" {
		color.New(color.Bold, color.FgYellow).Println("Failed to fetch Git global user email:", err.Error())
	}
	email := strings.TrimSpace(string(output))
	return email
}
