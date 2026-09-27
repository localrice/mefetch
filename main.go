package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/goccy/go-yaml"
)

func initConfig() error {
	configDir := os.Getenv("HOME") + "/.config/mefetch"

	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		return err
	}

	configPath := configDir + "/config.yaml"
	asciiPath := configDir + "/ascii.txt"

	if _, err := os.Stat(configPath); err == nil {
		return fmt.Errorf("mefetch is already initialized")
	}

	config := `name: Your Name
location: Your Location
interests: Your Interests

github: your-github
discord: your-discord

text-color: "#4ba3f5"
ascii: true
`

	ascii := `   /\_/\\
  ( o.o )
   > ^ <
`

	err = os.WriteFile(configPath, []byte(config), 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(asciiPath, []byte(ascii), 0644)
	if err != nil {
		return err
	}

	return nil
}

func openConfig() error {
	configPath := os.Getenv("HOME") + "/.config/mefetch/config.yaml"

	editor := os.Getenv("EDITOR")

	if editor == "" {
		editor = "nano"
	}

	cmd := exec.Command(editor, configPath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

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

// do not use anything other than color in the text-color field, as it will defualt to the usual white
// this function does not have any color validation
// only supports hex or ANSI256 value
func loadColor(config yaml.MapSlice) (string, error) {
	for _, item := range config {
		if item.Key == "text-color" {
			return fmt.Sprintf("%v", item.Value), nil
		}
	}
	return "#4ba3f5", nil
}

func main() {
	if len(os.Args) > 1 {

		if os.Args[1] == "init" {
			err := initConfig()

			if err != nil {
				fmt.Println("Error initializing mefetch:", err)
				return
			}

			fmt.Println("mefetch initialized successfully.")
			return
		}

		if os.Args[1] == "config" {
			err := openConfig()

			if err != nil {
				fmt.Println("Error opening config:", err)
			}

			return
		}
	}

	path := os.Getenv("HOME") + "/.config/mefetch/config.yaml"
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("mefetch is not initialized.")
			fmt.Println("run 'mefetch init' to create the default configuration")
		}
		fmt.Println("Error reading file", err)
		return
	}

	var config yaml.MapSlice

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling YAML", err)
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

	var content strings.Builder

	for _, item := range config {
		// to remove the ascii key from the output
		if item.Key == "ascii" || item.Key == "text-color" {
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
