package scheduler

import (
	"errors"
	"time"
)

// Schedule defines allowed activity hours
type Schedule struct {
	StartHour int // e.g., 9
	EndHour   int // e.g., 18
}

// NewDefaultSchedule returns business-hours schedule
func NewDefaultSchedule() Schedule {
	return Schedule{
		StartHour: 9,
		EndHour:   18,
	}
}

// Allow checks if current time is within allowed hours
func (s Schedule) Allow() error {
	now := time.Now()
	hour := now.Hour()

	if hour < s.StartHour || hour >= s.EndHour {
		return errors.New("outside allowed activity hours")
	}

	return nil
}
