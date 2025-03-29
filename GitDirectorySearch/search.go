package gitdirectorysearch

import (
	"os"
	"path/filepath"

	color "github.com/fatih/color"
)

type Search struct{}

func (gd Search) TrackGitDirs() []string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("No Home Directory Found ")
		return nil
	}

	var gitDirs []string

	err = filepath.WalkDir(userHomeDir, func(path string, d os.DirEntry, err error) error {
		if err != nil { // skip unwalkable directory
			return nil
		}

		if d.IsDir() && d.Name() == ".git" {
			gitDirs = append(gitDirs, path)
		}
		return nil
	})
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("Error occured while searching commits:", err)
	}
	return gitDirs
}
