package models

import (
	"time"
)

type State int

const (
	Unmarked = iota
	InProgress
	Completed
)

type Task struct {
	ID        int
	Text      string
	Cursor    int // Stores in which letter of the string the cursor was
	State     State
	StartTime time.Time
	EndTime   time.Time
}
