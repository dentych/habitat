package vscode

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	_ "embed"

	"gitlab.com/dentych/habitat/terminal"
)

//go:embed extensions.txt
var extensions string

//go:embed keybindings.json
var keybindings []byte

//go:embed settings.json
var settings []byte

func InstallVSCode(homeDir string) {
	settingsDir := getVsCodeSettingsDirectory(homeDir)
	fmt.Println("Creating VSCode settings directory if it doesn't exist: ", settingsDir)
	err := os.MkdirAll(settingsDir, 0755)
	if err != nil {
		log.Fatalln("There was an error making the VsCode settings directory: ", err)
	}

	keyBindingPath := getKeybindingPath(settingsDir)
	settingsPath := getSettingsPath(settingsDir)

	fmt.Println("Writing keybindings.json...")
	err = os.WriteFile(keyBindingPath, keybindings, 0644)
	if err != nil {
		log.Fatalln("Failed to write keybindings.json file: ", err)
	}
	fmt.Println("keybindings.json successfully written")

	fmt.Println("Writing settings.json...")
	err = os.WriteFile(settingsPath, settings, 0644)
	if err != nil {
		log.Fatalln("Failed to write settings.json file: ", err)
	}
	fmt.Println("settings.json successfully written")

	fmt.Println("Installing extensions...")
	installExtensions()
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

func installExtensions() {
	fmt.Println("Ensuring code command exists")
	cmd := exec.Command("code", "--version")
	err := cmd.Run()
	if err != nil {
		log.Fatalln("Code command not found: ", err)
	}

	extensionList := strings.Split(strings.TrimSpace(extensions), "\n")
	fmt.Println("Installing the following extensions:")
	for _, ext := range extensionList {
		if len(ext) > 0 {
			fmt.Println("- " + ext)
		}
	}

	for _, ext := range extensionList {
		if len(strings.TrimSpace(ext)) == 0 {
			continue
		}

		terminal.PrintHeading(fmt.Sprintf("Installing extension: %s", ext))
		cmd := exec.Command("code", "--install-extension", ext)
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatalln("Failed to install extensions: ", err)
		}
		fmt.Println()
		fmt.Println()
	}
}
