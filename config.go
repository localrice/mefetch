package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/goccy/go-yaml"
)

var (
	configDir  = os.Getenv("HOME") + "/.config/mefetch"
	configPath = configDir + "/config.yaml"
	asciiPath  = configDir + "/ascii.txt"
)

func initConfig() error {
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		return err
	}

	if _, err := os.Stat(configPath); err == nil {
		return fmt.Errorf("mefetch is already initialized")
	}
	config := `name: auto
location: auto
bio: auto
github: your-github
interests: cats

text-color: "#B46A72"
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
				data, err := os.ReadFile(asciiPath)
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
	// default color if nothing is set in the config.yaml file
	return "#4ba3f5", nil
}

func loadConfig() (yaml.MapSlice, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config yaml.MapSlice

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling YAML", err)
		return nil, err
	}
	return config, nil
}

func loadGitHubProfile(username string) (map[string]interface{}, error) {
	url := "https://api.github.com/users/" + username

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("GitHub username not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %s", resp.Status)
	}
	var github_profile map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&github_profile)
	if err != nil {
		return nil, err
	}

	return github_profile, nil
}
