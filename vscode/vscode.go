package vscode

import (
	"fmt"
	"log"
	"runtime"
)

func ScrapeVSCodeConfig(homeDir string) {
	keybindingPath := getKeybindingPath(homeDir)
	settingsPath := getSettingsPath(homeDir)

	if keybindingPath == "" || settingsPath == "" {
		log.Fatalf("Unsupported platform detected: %s\n", runtime.GOOS)
	}
}

func getKeybindingPath(homeDir string) string {
	switch runtime.GOOS {
	case "darwin":
		return fmt.Sprintf("%s/Library/Application Support/Code/User/keybindings.json", homeDir)
	case "windows":
		return fmt.Sprintf("%s/AppData/Roaming/Code/User/keybindings.json", homeDir)
	}

	return ""
}

func getSettingsPath(homeDir string) string {
	switch runtime.GOOS {
	case "darwin":
		return fmt.Sprintf("%s/Library/Application Support/Code/User/settings.json", homeDir)
	case "windows":
		return fmt.Sprintf("%s/AppData/Roaming/Code/User/keybindings.json", homeDir)
	}

	return ""
}
