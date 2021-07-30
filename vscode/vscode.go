package vscode

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func InstallVSCode(homeDir string) {
	fmt.Println("Creating VSCode settings directory if it doesn't exist: ", getVsCodeSettingsDirectory(homeDir))
	err := os.MkdirAll(getVsCodeSettingsDirectory(homeDir), 0755)
	if err != nil {
		log.Fatalln("There was an error making the VsCode settings directory: ", err)
	}
}

func ScrapeVSCode(homeDir string) {
	fmt.Println("Scraping VS Code configuration files")
	settingsDir := getVsCodeSettingsDirectory(homeDir)
	if settingsDir == "" {
		log.Fatalf("Unsupported platform detected: %s\n", runtime.GOOS)
	}
	keybindingPath := getKeybindingPath(settingsDir)
	settingsPath := getSettingsPath(settingsDir)

	keybindingJson, err := os.ReadFile(keybindingPath)
	if err != nil {
		log.Fatalln("Failed to read keybindings: ", err)
	}
	os.WriteFile("vscode/keybindings.json", keybindingJson, 0644)

	settingsJson, err := os.ReadFile(settingsPath)
	if err != nil {
		log.Fatalln("Failed to read settings: ", err)
	}
	os.WriteFile("vscode/settings.json", settingsJson, 0644)

	cmd := exec.Command("code", "--version")
	_, err = cmd.Output()
	if err != nil {
		log.Fatalln("Err command: ", err)
	}

	cmd = exec.Command("code", "--list-extensions")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalln("Err command: ", err)
	}

	os.WriteFile("vscode/extensions.txt", output, 0644)
}

func getVsCodeSettingsDirectory(homeDir string) string {
	switch runtime.GOOS {
	case "darwin":
		return fmt.Sprintf("%s/Library/Application Support/Code/User", homeDir)
	case "windows":
		return fmt.Sprintf("%s/AppData/Roaming/Code/User/keybindings.json", homeDir)
	}
	return ""
}
func getKeybindingPath(settingsDir string) string {
	return fmt.Sprintf("%s/keybindings.json", settingsDir)
}

func getSettingsPath(settingsDir string) string {
	return fmt.Sprintf("%s/settings.json", settingsDir)
}
