package menus

import (
	"gitlab.com/dentych/env/internal/modules"
	"gitlab.com/dentych/env/internal/terminal"
)

type BashMenu struct {
	DefaultMenu
	BashModule modules.Module
}

func NewBashMenu(parent Menu) *BashMenu {
	menu := BashMenu{}
	menu.Name = "Bash menu"
	menu.Parent = parent
	menu.BashModule = modules.NewBashModule()
	menu.Options = []Option{
		{Key: '1', Description: "Install", Handler: menu.install},
		{Key: '2', Description: "Uninstall", Handler: menu.uninstall},
		{Key: 'q', Description: "Back", Handler: func() Menu { return menu.Parent }},
	}

	return &menu
}

func (m *BashMenu) install() Menu {
	m.BashModule.Install()
	terminal.Read()
	return m
}

func (m *BashMenu) uninstall() Menu {
	m.BashModule.Uninstall()
	terminal.Read()
	return m
}
