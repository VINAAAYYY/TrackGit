package contributionvisualizer

import (
	models "TrackGit/DatabaseManager/Models"
	"fmt"
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (c ContributionVisualizer) BuildCalender(startDate time.Time) ContributionVisualizer {
	// Get commit counts and build a date-to-commit-count map.
	commits := c.GetCommitCount(startDate)
	commitDateMap := mapCommitToContributionDate(commits)

	c.table = tview.NewTable().
		SetFixed(2, 1).
		SetSelectable(true, true)

	// Calculate weeks since the start date.
	weeks := int(math.Ceil(float64(time.Since(startDate).Hours() / 24 / 7)))

	// Build header rows.
	buildMonthHeader(c.table, startDate, weeks)
	buildDayHeader(c.table)

	// Fill in the calendar cells (starting at row 2).
	fillCalendarCells(c.table, startDate, weeks, commitDateMap)

	// Initialize the application and tooltip.
	c.app, c.tooltip = getAppAndToolTip()
	// Set up cell click behavior.
	c.table.SetSelectedFunc(c.getOnClickFunc(commitDateMap))
	return c
}

func buildMonthHeader(table *tview.Table, startDate time.Time, weeks int) {
	var prevMonth time.Month
	for col := 1; col <= weeks; col++ {
		weekStartDate := startDate.AddDate(0, 0, (col-1)*7)
		currentMonth := weekStartDate.Month()
		// Show the month label if it's the first week or if the month changes.
		monthLabel := ""
		if col == 1 || currentMonth != prevMonth {
			monthLabel = currentMonth.String()
			prevMonth = currentMonth
		}
		table.SetCell(0, col,
			tview.NewTableCell(monthLabel).
				SetAlign(tview.AlignCenter).
				SetSelectable(false))
	}
}

func buildDayHeader(table *tview.Table) {
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	for row, day := range days {
		table.SetCell(row+2, 0,
			tview.NewTableCell(day).
				SetAlign(tview.AlignCenter).
				SetSelectable(false))
	}
}

func fillCalendarCells(table *tview.Table, startDate time.Time, weeks int, commitDateMap map[time.Time]int) {
	// Calculate how many cells to skip in the first week.
	firstWeekday := int(startDate.Weekday()) // 0=Sunday, 6=Saturday
	// Normalize the start date to midnight.
	date := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())

	totalCells := weeks * 7
	for i := 0; i < totalCells; i++ {
		col := (i / 7) + 1 // Calculate column index (starting at 1)
		row := i % 7       // Calculate row index (0-6)

		// For the first week, leave cells blank for days before startDate.
		if col == 1 && row < firstWeekday {
			table.SetCell(row+2, col,
				tview.NewTableCell("").
					SetSelectable(false))
			continue
		}

		commitCount := commitDateMap[date]
		cellColor := colorForCommitCount(commitCount)
		cell := tview.NewTableCell(fmt.Sprintf("%d", commitCount)).
			SetBackgroundColor(cellColor).
			SetAlign(tview.AlignLeft).
			SetReference(date)
		table.SetCell(row+2, col, cell)
		date = date.AddDate(0, 0, 1) // next date
	}
}

func getAppAndToolTip() (*tview.Application, *tview.Box) {
	app := tview.NewApplication()
	app.EnableMouse(true)
	tooltip := tview.NewTextView().
		SetBorder(true).
		SetTitle("TrackGit's Commit Visualizer").
		SetTitleColor(tcell.ColorLavender)

	return app, tooltip
}

func (c *ContributionVisualizer) getOnClickFunc(commitDateMap map[time.Time]int) func(row int, column int) {
	return func(row int, column int) {
		fmt.Println("On Click Called")

		cell := c.table.GetCell(row, column)
		if cell == nil || cell.GetReference() == nil {
			return
		}

		date := cell.GetReference().(time.Time) // Convert reference to date
		commitCount := commitDateMap[date]

		// Update tooltip text
		c.tooltip.SetTitle(fmt.Sprintf(" Date: %s  Commits: %d ", date.Format("2006-01-02"), commitCount))

		// Sync app to refresh screen and remove deprecated screen elements
		c.app.Sync()
	}
}

func mapCommitToContributionDate(commits []models.CommitHistory) map[time.Time]int {
	// Get the user's current timezone
	userLocation, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}

	commitMap := make(map[time.Time]int)
	for _, commit := range commits {
		// Ensure the commit date is in UTC before converting to the user's timezone
		utcTime := commit.Date.UTC()
		// Convert the commit date from UTC to the user's timezone and extract the date part
		localTime := utcTime.In(userLocation)
		dateOnly := time.Date(localTime.Year(), localTime.Month(), localTime.Day(), 0, 0, 0, 0, localTime.Location())
		commitMap[dateOnly]++
	}
	return commitMap
}

func colorForCommitCount(count int) tcell.Color {
	switch {
	case count == 0:
		return tcell.ColorGray
	case count < 5:
		return tcell.ColorLightGreen
	case count < 10:
		return tcell.NewRGBColor(0, 150, 0) // dark green
	default:
		return tcell.NewRGBColor(0, 100, 0) // light gray
	}
}
