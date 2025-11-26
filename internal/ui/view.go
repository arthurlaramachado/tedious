package ui

import (
	"fmt"

	comp "github.com/arthurlaramachado/tedious/internal/components"
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
				line = comp.BlueLine.Render(line)
			} else {
				line = comp.GrayLine.Render(line)
			}
		case InProgress:
			line = comp.YellowLine.Render(line)
		case Completed:
			line = comp.GreenLine.Render(line)
		}

		s += line + "\n"
	}
	s += fmt.Sprintf("\nTask counter: %d\nUse ↑/↓, enter to toggle, ctrl+c to quit.", m.taskCounter)
	return s
}
