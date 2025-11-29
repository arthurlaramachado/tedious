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

func GetInfoLine(width int, isSubtask bool) lipgloss.Style {
	var paddingLeft int

	if isSubtask {
		paddingLeft = 8
	} else {
		paddingLeft = 4
	}

	return lipgloss.NewStyle().
		Italic(true).
		PaddingLeft(paddingLeft).
		Foreground(lipgloss.Color("#ffffffff")).
		Bold(true).
		Width(width - padding)
}

func GetLineColor(c Color, width int, isSubtask bool) lipgloss.Style {
	var color_hex string
	var paddingLeft int = 0

	if isSubtask {
		paddingLeft = 4
	}

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
		PaddingLeft(paddingLeft).
		Width(width - padding)
}
