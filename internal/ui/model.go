package ui

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

type Model struct {
	tasks       []Task
	cursor      int // Stores in which task we're in
	taskCounter int
}

func NewModel() Model {
	var m = Model{
		tasks:       []Task{},
		cursor:      0,
		taskCounter: 0,
	}

	m.CreateEmptyTask()

	return m
}

func (m *Model) CreateEmptyTask() {
	m.tasks = append(m.tasks, Task{
		ID:        m.taskCounter,
		Text:      "",
		State:     Unmarked,
		StartTime: time.Time{},
		EndTime:   time.Time{},
	})

	m.taskCounter++
}

func (m *Model) RemoveTextLastLetter() {
	task := &m.tasks[m.cursor]
	if len(task.Text) > 0 {
		task.Text = task.Text[:len(task.Text)-1]
	}
}

func (m *Model) DeleteTask() {
	m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)
}

func (m *Model) AddChar(char string) {
	m.tasks[m.cursor].Text += char
}
