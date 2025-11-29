package ui

import (
	"time"

	mod "github.com/arthurlaramachado/tedious/internal/models"
	"github.com/arthurlaramachado/tedious/internal/utils"
)

type Model struct {
	tasks       []mod.Task
	cursor      int // Stores in which task we're in
	taskCounter int
	width       int
	height      int
}

func StartModel() Model {
	var tasks []mod.Task
	jsonBlob := utils.ReadState()

	if jsonBlob != nil {
		tasks = jsonBlob
	} else {
		tasks = []mod.Task{}
	}

	var m = Model{
		tasks:       tasks,
		cursor:      0,
		taskCounter: len(tasks),
	}

	if len(tasks) == 0 {
		m.CreateEmptyTask()
	}

	return m
}

func (m *Model) RearrangeTasksUp(index int) {
	copy := m.tasks[index-1]
	m.tasks[index-1] = m.tasks[index]
	m.tasks[index] = copy
}

func (m *Model) RearrangeTasksDown(index int) {
	copy := m.tasks[index+1]
	m.tasks[index+1] = m.tasks[index]
	m.tasks[index] = copy
}

func (m *Model) CreateEmptyTask() {
	m.tasks = append(m.tasks, mod.Task{
		ID:        m.taskCounter,
		Text:      "",
		State:     mod.Unmarked,
		StartTime: time.Time{},
		EndTime:   time.Time{},
	})

	m.taskCounter++
}

func (m *Model) CreateEmptyTaskAt(index int) {
	newTask := mod.Task{
		ID:        m.taskCounter,
		Text:      "",
		State:     mod.Unmarked,
		StartTime: time.Time{},
		EndTime:   time.Time{},
	}

	m.tasks = append(m.tasks[:index], append([]mod.Task{newTask}, m.tasks[index:]...)...)

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
