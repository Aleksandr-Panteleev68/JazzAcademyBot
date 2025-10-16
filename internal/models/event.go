package models

import "time"

type Event struct {
	ID          string    `db:"id"` // Google Calendare event ID
	Title       string    `db:"title"`
	Description string    `db:"descrioption"` // Raw JSON from Google
	Deadline    time.Time `db:"deadline"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type EventDescription struct {
	RemindersInternal string `json:"reminders_internal"` // e.g., "1h"
	RethinkPeriod     string `json:"rethink_period"`     // e.g., "24h" for "подумают" reminders
}
