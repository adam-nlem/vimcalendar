package models

import "time"

type Event struct {
	ID            int
	Title         string
	StartDateTime time.Time
	EndDateTime   time.Time
}
