package vscode

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func ScrapeVSCodeConfig(homeDir string) {
	fmt.Println("Scraping VS Code configuration files")
	keybindingPath := getKeybindingPath(homeDir)
	settingsPath := getSettingsPath(homeDir)

	if keybindingPath == "" || settingsPath == "" {
		log.Fatalf("Unsupported platform detected: %s\n", runtime.GOOS)
	}

	keybindingJson, err := os.ReadFile(keybindingPath)
	if err != nil {
		log.Fatal("Failed to read keybindings:", err)
	}
	os.WriteFile("keybindings.json", keybindingJson, 0644)

	settingsJson, err := os.ReadFile(settingsPath)
	if err != nil {
		log.Fatal("Failed to read settings:", err)
	}
	os.WriteFile("settings.json", settingsJson, 0644)
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
