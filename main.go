package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// STYLES
var greenLine = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#04B575"))

var blueLine = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4a77ffff"))

var yellowLine = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#e5de00ff"))

var grayLine = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#9a9a9aff"))

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

func removeLastLetter(text *string) {
	if len(*text) > 0 {
		*text = (*text)[:len(*text)-1]
	}
}

func deleteTask(tasks *[]Task, cursor *int) {
	*tasks = append((*tasks)[:(*cursor)], (*tasks)[(*cursor)+1:]...)
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
				m.cursor++
			}

		case "backspace":
			lineText := &m.tasks[m.cursor].Text

			if len(*lineText) == 0 && m.cursor > 1 {
				deleteTask(&m.tasks, &m.cursor)
			} else {
				removeLastLetter(lineText)
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

		if msg.Type == tea.KeyRunes {
			m.tasks[m.cursor].Text += string(msg.Runes)
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Tasks:\n"
	for i, task := range m.tasks {
		if i == m.cursor {
			s += ">"
		} else {
			s += " "
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

		line := fmt.Sprintf("[%s] %s", state, task.Text)

		switch task.State {
		case Unmarked:
			if i == m.cursor {
				line = blueLine.Render(line)
			} else {
				line = grayLine.Render(line)
			}
		case InProgress:
			line = yellowLine.Render(line)
		case Completed:
			line = greenLine.Render(line)
		}

		s += line + "\n"
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

// TODOs
/*
	Refactor to organize code
	Make it adapt to the whole terminal size
	clear terminal, only showing list
	Fix words to break at the end of the line
	Add persistency
	Add ways to press "tab" to create subitens
*/
