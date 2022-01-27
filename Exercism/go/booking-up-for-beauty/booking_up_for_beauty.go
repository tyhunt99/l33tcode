package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const timeLayout = "1/2/2006 15:04:05"
	parsedTime, err := time.Parse(timeLayout, date)
	if err != nil {
		panic(err)
	}
	return parsedTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const timeLayout = "January 2, 2006 15:04:05"
	parsedTime, err := time.Parse(timeLayout, date)
	if err != nil {
		panic(err)
	}
	return !parsedTime.After(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const timeLayout = "Monday, January 2, 2006 15:04:05"
	parsedTime, err := time.Parse(timeLayout, date)
	if err != nil {
		panic(err)
	}
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const timeLayout = "Monday, January 2, 2006, at 15:04"
	appt := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s.", appt.Format(timeLayout))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
