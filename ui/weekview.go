package ui

import (
	"github.com/rivo/tview"

	"github.com/adam-nlem/vimcalendar/services"
)

const WeekPageName = "week"

func WeekPage() *tview.Box {
	view := tview.NewBox().SetBorder(true).SetTitle("Week View")

	view.SetInputCapture(services.WeekViewInputCapture)

	return view
}
