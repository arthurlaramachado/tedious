package ui

import (
	"fmt"

	comp "github.com/arthurlaramachado/tedious/internal/components"
	mod "github.com/arthurlaramachado/tedious/internal/models"
)

func (m Model) View() string {
	s := "Tasks:\n"
	for i, task := range m.tasks {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
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
				line = comp.GetLineColor(comp.BLUE, m.width, task.IsSubtask).Render(line)
			} else {
				line = comp.GetLineColor(comp.GREY, m.width, task.IsSubtask).Render(line)
			}
		case mod.InProgress:
			line = comp.GetLineColor(comp.YELLOW, m.width, task.IsSubtask).Render(line)
		case mod.Completed:
			line = comp.GetLineColor(comp.GREEN, m.width, task.IsSubtask).Render(line)

			duration := task.EndTime.Sub(task.StartTime)

			hours := int(duration.Hours())
			minutes := int(duration.Minutes()) % 60
			seconds := int(duration.Seconds()) % 60

			timeInfo := fmt.Sprintf("Duration: %dh %dm %ds", hours, minutes, seconds)
			line += "\n" + comp.GetInfoLine(m.width, task.IsSubtask).Render(timeInfo)
		}

		s += cursor + line + "\n"
	}
	s += fmt.Sprintf("\nTask counter: %d\nUse ↑/↓, enter to toggle, ctrl+c to quit.", m.taskCounter)
	return s
}
