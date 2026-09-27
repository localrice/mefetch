package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/goccy/go-yaml"
)

func main() {
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

	// styling definitions
	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4ba3f5")).
		Bold(true)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)

	var content strings.Builder
	var asciiArt string

	for _, item := range config {
		if item.Key == "ascii" {
			if item.Value == true {
				path := os.Getenv("HOME") + "/.config/mefetch/ascii.txt"

				data, err := os.ReadFile(path)
				if err != nil {
					fmt.Println("Error reading ascii.txt", err)
					return
				}

				asciiArt = string(data)
			}

			continue
		}

		label := fmt.Sprintf("%-10s", item.Key)
		styledLabel := labelStyle.Render(label)

		fmt.Fprintf(&content, "%s > %v\n", styledLabel, item.Value)
	}

	fmt.Println(asciiArt)
	fmt.Println(boxStyle.Render(content.String()))
}
