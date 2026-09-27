package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/goccy/go-yaml"
)

func loadASCII(config yaml.MapSlice) (string, error) {
	for _, item := range config {
		if item.Key == "ascii" {
			if item.Value == true {
				path := os.Getenv("HOME") + "/.config/mefetch/ascii.txt"

				data, err := os.ReadFile(path)
				if err != nil {
					return "", err
				}

				return string(data), nil
			}

			return "", nil
		}

	}
	return "", nil
}

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
		Padding(1, 1)

	asciiStyle := lipgloss.NewStyle().
		PaddingRight(5)

	asciiArt, err := loadASCII(config)
	if err != nil {
		fmt.Println("Error reading ascii.txt:", err)
		return
	}

	ascii := asciiStyle.Render(asciiArt)

	var content strings.Builder

	for _, item := range config {
		if item.Key == "ascii" {
			continue
		}

		label := fmt.Sprintf("%-10s", item.Key)
		styledLabel := labelStyle.Render(label)

		fmt.Fprintf(&content, "%s > %v\n", styledLabel, item.Value)
	}

	// fmt.Println(asciiArt)
	// fmt.Println(boxStyle.Render(content.String()))
	profile := content.String()

	output := lipgloss.JoinHorizontal(
		lipgloss.Top,
		ascii,
		profile,
	)

	output = boxStyle.Render(output)

	fmt.Println(output)

}
