package main

import (
	"log"
	"os"
	"time"
	ContributionVisualizer "TrackGit/ContributionVisualizer"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No Time Frame provided")
	}

	timeFrame := os.Args[1]
	startDate := getStartDate(timeFrame)
	var c ContributionVisualizer.ContributionVisualizer
	c.GetCalenderFlex(startDate)
}

func getStartDate(timeFrame string) time.Time {
	if timeFrame == "week" {
		return time.Now().AddDate(0, 0, -7)
	} else if timeFrame == "month" {
		return time.Now().AddDate(0, -1, 0)
	} else if timeFrame == "year" {
		return time.Now().AddDate(-1, 0, 0)
	}
	return time.Now().AddDate(0, 0, -7)
}
