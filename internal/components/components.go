package components

import (
	"github.com/charmbracelet/lipgloss"
)

type Color int

const (
	GREY Color = iota
	YELLOW
	BLUE
	GREEN
)

const padding = 4

func GetInfoLine(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Italic(true).
		PaddingLeft(4).
		Foreground(lipgloss.Color("#ffffffff")).
		Bold(true).
		Width(width - padding)
}

func GetLineColor(c Color, width int) lipgloss.Style {
	var color_hex string

	switch c {
	case GREY:
		color_hex = "#9a9a9aff"
	case YELLOW:
		color_hex = "#e5de00ff"
	case BLUE:
		color_hex = "#4a77ffff"
	case GREEN:
		color_hex = "#04B575"
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(color_hex)).
		Width(width - padding)
}
