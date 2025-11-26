package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.CreateEmptyTask()
				m.cursor++
			}

		case "backspace":
			lineText := &m.tasks[m.cursor].Text

			if len(*lineText) == 0 && m.cursor > 1 {
				m.DeleteTask()
			} else {
				m.RemoveTextLastLetter()
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
			m.AddChar(string(msg.Runes))
		}
	}

	return m, nil
}
