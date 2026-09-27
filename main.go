package main

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/goccy/go-yaml"
)

func main() {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#4ba3f5")).
		Padding(2).
		PaddingLeft(4).
		Width(22)

	lipgloss.Println(style.Render(" welcome to ricefield org"))

	path := os.Getenv("HOME") + "/.config/mefetch/config.yaml"
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var config yaml.MapSlice

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling YAML", err)
		return
	}

	for i, item := range config {
		fmt.Printf("%d. %v >  %v\n", i+1, item.Key, item.Value)
	}

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4ba3f5")).
		Bold(true)

	for _, item := range config {
		fmt.Printf(
			"%s > %v\n",
			labelStyle.Render(fmt.Sprintf("%-10s", item.Key)),
			item.Value,
		)
	}
}
