package contributionvisualizer

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

func (c ContributionVisualizer) GetCalenderFlex(startDate time.Time) *tview.Flex {
	calender := c.BuildCalender(startDate)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(calender.table, 0, 1, false).
		AddItem(calender.tooltip, 3, 1, false)
	
	if calender.app != nil {
		if err := calender.app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	}

	return flex
}

func printCalender(calendarTable *tview.Table) {
	if calendarTable == nil {
		fmt.Println("calendarTable is nil")
	} else {
		fmt.Println("calendarTable initialized")
		hasContent := false
		for row := 0; row < 10; row++ { // Adjust the row limit based on expected size
			for col := 0; col < 10; col++ { // Adjust the col limit based on expected size
				cell := calendarTable.GetCell(row, col)
				if cell != nil {
					text := cell.Text
					fmt.Printf("Cell [%d, %d]: %s\n", row, col, text)
					hasContent = true
				}
			}
		}
		if !hasContent {
			fmt.Println("calendarTable is empty")
		}
	}
}
