package main

import "charm.land/lipgloss/v2"

func main() {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#4ba3f5")).
		Padding(2).
		PaddingLeft(4).
		Width(22)

	lipgloss.Println(style.Render(" welcome to ricefield org"))
}
