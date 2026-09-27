package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/goccy/go-yaml"
)

func main() {
	// style := lipgloss.NewStyle().
	// 	Bold(true).
	// 	Foreground(lipgloss.Color("#FAFAFA")).
	// 	Background(lipgloss.Color("#4ba3f5")).
	// 	Padding(2).
	// 	PaddingLeft(4).
	// 	Width(22)

	// lipgloss.Println(style.Render(" welcome to ricefield org"))

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

	// for i, item := range config {
	// 	fmt.Printf("%d. %v >  %v\n", i+1, item.Key, item.Value)
	// }

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4ba3f5")).
		Bold(true)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)

	var content strings.Builder

	for _, item := range config {
		label := fmt.Sprintf("%-10s", item.Key)
		styledLabel := labelStyle.Render(label)

		fmt.Fprintf(&content, "%s > %v\n", styledLabel, item.Value)
	}
	fmt.Println(boxStyle.Render(content.String()))
}
