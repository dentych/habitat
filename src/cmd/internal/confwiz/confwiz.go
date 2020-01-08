package confwiz

import (
	"bufio"
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
	moduleEnabled("Git", scanner, &w.configuration.ModulesEnabled.Git)
	moduleEnabled("Tmux", scanner, &w.configuration.ModulesEnabled.Tmux)
	moduleEnabled("Vim", scanner, &w.configuration.ModulesEnabled.Vim)
	moduleEnabled("Bash", scanner, &w.configuration.ModulesEnabled.Bash)

	fmt.Printf("### Git ###\n")
	fmt.Printf("Name (%s): ", w.configuration.Git.Name)
	if scanner.Scan() {
		scannerOutput = strings.Trim(scanner.Text(), " ")
		if scannerOutput != "" {
			w.configuration.Git.Name = scannerOutput
		}
	}

	fmt.Printf("Email (%s): ", w.configuration.Git.Email)
	if scanner.Scan() {
		scannerOutput = strings.Trim(scanner.Text(), " ")
		if scannerOutput != "" {
			w.configuration.Git.Email = scannerOutput
		}
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

func moduleEnabled(name string, scanner *bufio.Scanner, element *bool) {
	scannerOutput := ""
	fmt.Printf("%s (%s): ", name, bla(*element))
	if scanner.Scan() {
		scannerOutput = strings.Trim(scanner.Text(), " ")
		if scannerOutput != "" {
			*element = parseYesNo(scannerOutput)
		}
	}
}
