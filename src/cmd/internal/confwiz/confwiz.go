package confwiz

import (
	"bufio"
	"bytes"
	"env/cmd/internal"
	"fmt"
	"os"
	"strings"
)

type ConfigurationWizard struct {
	configuration internal.Configuration
}

func New(configuration *internal.Configuration) *ConfigurationWizard {
	if configuration == nil {
		configuration = internal.NewDefaultConfiguration()
	}
	return &ConfigurationWizard{
		configuration: *configuration,
	}
}

func (w *ConfigurationWizard) Configure() internal.Configuration {
	fmt.Printf("Welcome to the ENV configuration wizard.\n\n")
	fmt.Printf("You will be asked for a variety of different options.\n")
	fmt.Printf("Default values are written in parentheses. If you wish to use the default, simply press enter.\n\n")

	var scannerOutput string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("### Modules enabled (default marked with caps) ###\n")
	fmt.Printf("Git (%s): ", bla(w.configuration.ModulesEnabled.Git))
	if scanner.Scan() {
		scannerOutput = strings.Trim(scanner.Text(), " ")
		w.configuration.ModulesEnabled.Git = parseYesNo(scannerOutput)
	}

	fmt.Printf("Tmux (%s): ", bla(w.configuration.ModulesEnabled.Tmux))
	if scanner.Scan() {
		scannerOutput = strings.Trim(scanner.Text(), " ")
		w.configuration.ModulesEnabled.Tmux = parseYesNo(scannerOutput)
	}

	fmt.Printf("Vim (%s): ", bla(w.configuration.ModulesEnabled.Vim))
	if scanner.Scan() {
		scannerOutput = strings.Trim(scanner.Text(), " ")
		w.configuration.ModulesEnabled.Vim = parseYesNo(scannerOutput)
	}

	fmt.Printf("### Git ###\n")
	fmt.Printf("Name (%s): ", w.configuration.Git.Name)
	if scanner.Scan() {
		scannerOutput = strings.Trim(scanner.Text(), " ")
		w.configuration.Git.Name = scannerOutput
	}

	fmt.Printf("Email (%s): ", w.configuration.Git.Email)
	if scanner.Scan() {
		scannerOutput = strings.Trim(scanner.Text(), " ")
		w.configuration.Git.Email = scannerOutput
	}

	fmt.Println()

	return w.configuration
}

func bla(value bool) string {
	if value {
		return "Y/n"
	}
	return "y/N"
}

func parseYesNo(value string) bool {
	if strings.ToLower(value) == "y" {
		return true
	}
	return false
}

func cleanText(text string) string {
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@.-_"
	var output bytes.Buffer
	for _, v := range text {
		if strings.ContainsRune(validChars, v) {
			output.WriteRune(v)
		}
	}

	return output.String()
}
