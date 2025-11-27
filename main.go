package main

import (
	"fmt"

	"github.com/arthurlaramachado/tedious/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := ui.StartModel()
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
