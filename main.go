package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
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

type model struct {
	tasks       []Task
	cursor      int // Stores in which task we're in
	taskCounter int
}

func initialModel() model {
	var m = model{
		tasks:       []Task{},
		cursor:      0,
		taskCounter: 0,
	}

	createEmptyTask(&m)

	return m
}

func createEmptyTask(m *model) {
	m.tasks = append(m.tasks, Task{
		ID:        m.taskCounter,
		Text:      "",
		State:     Unmarked,
		StartTime: time.Time{},
		EndTime:   time.Time{},
	})
	m.cursor++
	m.taskCounter++
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			} else if m.cursor == len(m.tasks)-1 && len(m.tasks[m.cursor].Text) != 0 {
				createEmptyTask(&m)
			}
		case "enter":
			task := &m.tasks[m.cursor]

			switch task.State {
			case Unmarked:
				task.State = InProgress
				task.StartTime = time.Now()
			case InProgress:
				task.State = Completed
				task.EndTime = time.Now()
			case Completed:
				task.State = Unmarked
				task.StartTime = time.Time{}
				task.EndTime = time.Time{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	s := "Tasks:\n"
	for i, task := range m.tasks {
		marker := " "
		if i == m.cursor {
			marker = ">"
		}
		state := " "
		switch task.State {
		case Unmarked:
			state = " "
		case InProgress:
			state = "~"
		case Completed:
			state = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", marker, state, task.Text)
	}
	s += fmt.Sprintf("\nTask counter: %d\nUse ↑/↓, enter to toggle, ctrl+c to quit.", m.taskCounter)
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
