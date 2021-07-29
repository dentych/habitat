package menus

import (
	"fmt"
	"gitlab.com/dentych/habitat/configuration"
	"gitlab.com/dentych/habitat/terminal"
)

type EditConfigMenu struct {
	DefaultMenu
}

func NewEditConfigMenu(parent Menu) *EditConfigMenu {
	menu := EditConfigMenu{}
	menu.Name = "Edit configuration menu"
	menu.Parent = parent
	menu.Options = []Option{
		{Key: '1', Description: "Git - Name", Handler: func() Menu { return menu.editConfigValue(&configuration.Config.Git.Name) }},
		{Key: '2', Description: "Git - Email", Handler: func() Menu { return menu.editConfigValue(&configuration.Config.Git.Email) }},
		{Key: '3', Description: "Git - Directory", Handler: func() Menu { return menu.editConfigValue(&configuration.Config.Git.Directory) }},
		{Key: 'q', Description: "Back", Handler: menu.Back},
	}

	return &menu
}

func (m *EditConfigMenu) editConfigValue(value *string) Menu {
	terminal.Clear()
	fmt.Println("Enter new value")
	input := terminal.Read()
	*value = input
	configuration.Config.Save()

	return m
}
