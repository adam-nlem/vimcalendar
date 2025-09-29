package services

import "time"

func GetTodaysDate() time.Time {
	now := time.Now()

	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func UpdateDate(date time.Time, add bool, numberOfDays int) time.Time {
	duration := 24 * time.Hour * time.Duration(numberOfDays)

	if add {
		duration = duration * -1
	}

	return date.Add(duration)
}

func GetTodaysWeekDates() []time.Time {
}

func UpdateWeek(weekDates []time.Time, add bool, numberOfWeeks int) []time.Time {
}
