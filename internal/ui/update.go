package ui

import (
	"time"

	mod "github.com/arthurlaramachado/tedious/internal/models"
	utils "github.com/arthurlaramachado/tedious/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			utils.SaveState(m.tasks)

			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "enter":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			} else if m.cursor == len(m.tasks)-1 && len(m.tasks[m.cursor].Text) != 0 {
				m.CreateEmptyTask()
				m.cursor++
			}

		case "ctrl+n":
			m.CreateEmptyTaskAt(m.cursor)

		case "ctrl+d":
			if len(m.tasks) > 1 {
				m.DeleteTask()

				if m.cursor == 1 {
					m.cursor--
				}
			}

		case "backspace":
			lineText := &m.tasks[m.cursor].Text

			if len(*lineText) == 0 && len(m.tasks) > 1 {
				m.DeleteTask()

				if m.cursor == 1 {
					m.cursor--
				}
			} else {
				m.RemoveTextLastLetter()
			}

		case "tab":
			task := &m.tasks[m.cursor]

			switch task.State {
			case mod.Unmarked:
				task.State = mod.InProgress
				task.StartTime = time.Now()
			case mod.InProgress:
				task.State = mod.Completed
				task.EndTime = time.Now()
			case mod.Completed:
				task.State = mod.Unmarked
				task.StartTime = time.Time{}
				task.EndTime = time.Time{}
			}

		case " ":
			m.AddChar(" ")
		}

		if msg.Type == tea.KeyRunes {
			m.AddChar(string(msg.Runes))
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}
