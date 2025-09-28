package ui

import (
	"github.com/adam-nlem/vimcalendar/services"
	"github.com/rivo/tview"
)

const DayPageName = "day"

func DayPage() *tview.Box {
	view := tview.NewBox().SetBorder(true).SetTitle("Day View")

	view.SetInputCapture(services.WeekViewInputCapture)

	return view
}
