package menus

import (
	"fmt"
	"gitlab.com/dentych/env/configuration"
	"gitlab.com/dentych/env/terminal"
)

type ConfigMenu struct {
	DefaultMenu
}

func NewConfigMenu(parent Menu) *ConfigMenu {
	menu := ConfigMenu{}
	menu.Name = "Configuration menu"
	menu.Parent = parent
	menu.Options = []Option{
		{Key: '1', Description: "List all configuration values", Handler: menu.listAllConfigValues},
		{Key: '2', Description: "Edit config values", Handler: menu.editConfigValues},
		{Key: 'q', Description: "Back", Handler: func() Menu {
			return parent
		}},
	}
	return &menu
}

func (m *ConfigMenu) listAllConfigValues() Menu {
	terminal.Clear()
	fmt.Println("Configuration items:")
	fmt.Println("--- GIT ---")
	fmt.Printf("Name: %s\n", valueOrDefault(configuration.Config.Git.Name))
	fmt.Printf("Email: %s\n", valueOrDefault(configuration.Config.Git.Email))
	fmt.Printf("Directory: %s\n", valueOrDefault(configuration.Config.Git.Directory))
	fmt.Println()
	fmt.Println("<enter> to go back")
	fmt.Println()
	terminal.Read()
	return m
}

func (m *ConfigMenu) editConfigValues() Menu {
	return NewEditConfigMenu(m)
}

func valueOrDefault(value string) string {
	if len(value) == 0 {
		return "<empty>"
	}
	return value
}