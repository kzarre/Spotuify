package main

import "github.com/charmbracelet/lipgloss"

// Colors
var (
	colBlack    = lipgloss.Color("#000000")
	colMagenta  = lipgloss.Color("#ff00ff")
	colGreen    = lipgloss.Color("#00ff41")
	colGreenDim = lipgloss.Color("#007a20")
	colRed      = lipgloss.Color("#ff3333")
	colWhite    = lipgloss.Color("#ffffff")
	colGray     = lipgloss.Color("#555555")
)

// Shared Layout Wrapper
func panel(content string, w int) string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(colMagenta).
		Background(colBlack).
		Width(w).
		Render(content)
}
