package gitdirectorysearch

import (
	"os"
	"strings"

	color "github.com/fatih/color"
)

type Search struct{}

func (gd Search) recursiveSearch(name string, gitInitiatedDirs *[]string) {
	if strings.HasSuffix(name, ".git/") {
		*gitInitiatedDirs = append(*gitInitiatedDirs, name)
	}
	files, err := os.ReadDir(name)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			nextPath := name + file.Name() + "/"
			gd.recursiveSearch(nextPath, gitInitiatedDirs)
		}
	}
}

func (gd Search) TrackGirDirs() []string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		color.New(color.Bold, color.FgRed).Println("No Home Directory Found")
	}

	gitInitiatedDirs := []string{}
	gd.recursiveSearch(userHomeDir+"/", &gitInitiatedDirs)
	return gitInitiatedDirs
}
