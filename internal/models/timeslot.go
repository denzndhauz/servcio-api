package models

import (
	"time"
)

type TimeSlot struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	IsAvailable bool      `json:"isAvailable"`
}
