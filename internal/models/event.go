package main

import (
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	TimeZone    string    `json:"time_zone"`
	Location    string    `json:"location"`
	IsAllDay    bool      `json:"is_all_day"`
	Recurrence  string    `json:"recurrence"`
}
