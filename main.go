package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
)

func main() {
	// cli handling
	if len(os.Args) > 1 {
		// mefetch init
		if os.Args[1] == "init" {
			err := initConfig()

			if err != nil {
				fmt.Println("Error initializing mefetch:", err)
				return
			}

			fmt.Println("mefetch initialized successfully.")
			return
		}
		// mefetch config
		// opens ~/.config/mefetch/config.yaml file in text editor
		if os.Args[1] == "config" {
			err := openConfig()

			if err != nil {
				fmt.Println("Error opening config:", err)
			}

			return
		}
	}

	config, err := loadConfig()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("mefetch is not initialized.")
			fmt.Println("run 'mefetch init' to create the default configuration")
			return
		}
		fmt.Println("Error reading file", err)
		return
	}

	var githubProfile map[string]interface{}
	githubConfigured := false

	for _, item := range config {
		if item.Key == "github" {
			username := fmt.Sprintf("%v", item.Value)
			if username == "" {
				fmt.Println("GitHub username is not set in config.yaml")
				return
			}
			githubConfigured = true
			githubProfile, err = loadGitHubProfile(username)
			if err != nil {
				fmt.Println("Error loading GitHub profile:", err)
				return
			}

			break
		}
	}

	if !githubConfigured {
		fmt.Println("GitHub username is not set in config.yaml")
		return
	}

	color, err := loadColor(config)
	if err != nil {
		fmt.Println("Error loading color:", err)
		return
	}

	// styling definitions
	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Bold(true)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 1)

	asciiStyle := lipgloss.NewStyle().
		PaddingRight(5)

	// load ascii art from ascii.txt if the ascii key is set to true in config.yaml
	asciiArt, err := loadASCII(config)
	if err != nil {
		fmt.Println("Error reading ascii.txt:", err)
		return
	}

	ascii := asciiStyle.Render(asciiArt)

	// profile extraction loop
	var content strings.Builder

	for _, item := range config {
		// to remove the ascii key from the output
		if item.Key == "ascii" || item.Key == "text-color" {
			continue
		}

		value := item.Value
		// we had use interface so it needs a string specficially
		key := item.Key.(string)
		if value == "auto" {
			githubValue, exists := githubProfile[key]

			if !exists || githubValue == nil {
				continue
			}

			value = githubValue
		}
		label := fmt.Sprintf("%-10s", item.Key)
		styledLabel := labelStyle.Render(label)

		fmt.Fprintf(&content, "%s > %v\n", styledLabel, value)
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
