package main

import (
	"fmt"

	"github.com/arthurlaramachado/tedious/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := ui.NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
