package ui

import (
	"fmt"

	comp "github.com/arthurlaramachado/tedious/internal/components"
	mod "github.com/arthurlaramachado/tedious/internal/models"
)

func (m Model) View() string {
	s := "Tasks:\n"
	for i, task := range m.tasks {
		if i == m.cursor {
			s += ">"
		} else {
			s += " "
		}
		state := " "
		switch task.State {
		case mod.Unmarked:
			state = " "
		case mod.InProgress:
			state = "~"
		case mod.Completed:
			state = "x"
		}

		line := fmt.Sprintf("[%s] %s", state, task.Text)

		switch task.State {
		case mod.Unmarked:
			if i == m.cursor {
				line = comp.BlueLine.Render(line)
			} else {
				line = comp.GrayLine.Render(line)
			}
		case mod.InProgress:
			line = comp.YellowLine.Render(line)
		case mod.Completed:
			line = comp.GreenLine.Render(line)
		}

		s += line + "\n"
	}
	s += fmt.Sprintf("\nTask counter: %d\nUse ↑/↓, enter to toggle, ctrl+c to quit.", m.taskCounter)
	return s
}
