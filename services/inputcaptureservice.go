package services

import "github.com/gdamore/tcell/v2"

// Custom inputs for week view
func WeekViewInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'l':
		// Go to next week
		return nil
	case 'h':
		// Go to previous week
		return nil
	}
	return event
}

// Custom inputs for week view
func DayViewInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'l':
		// Go to next day
		return nil
	case 'h':
		// Go to previous day
		return nil
	}
	return event
}
