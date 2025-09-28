package main

import (
	"github.com/adam-nlem/vimcalendar/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	pages := tview.NewPages()

	pages.AddPage(ui.WeekPageName, ui.WeekPage(), true, true)
	pages.AddPage(ui.DayPageName, ui.DayPage(), true, false) // Hidden page

	app := tview.NewApplication()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'v':
			if currentPageName, _ := pages.GetFrontPage(); currentPageName == ui.DayPageName {
				pages.SwitchToPage(ui.WeekPageName)
			} else {
				pages.SwitchToPage(ui.DayPageName)
			}
			return nil
		case 'q':
			app.Stop()
			return nil
		}

		return event
	})
	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
