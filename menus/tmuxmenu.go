package menus

import (
	"gitlab.com/dentych/env/modules"
	"gitlab.com/dentych/env/terminal"
)

type TmuxMenu struct {
	DefaultMenu
}

func NewTmuxMenu(parent Menu) *TmuxMenu {
	menu := TmuxMenu{}
	menu.Name = "Tmux menu"
	menu.Parent = parent
	menu.Options = []Option{
		{Key: '1', Description: "Install", Handler: menu.install},
		{Key: '2', Description: "Uninstall", Handler: menu.uninstall},
		{Key: 'q', Description: "Back", Handler: menu.backHandler},
	}
	menu.Module = modules.Tmux{}

	return &menu
}

func (m *TmuxMenu) install() Menu {
	m.Module.Install()
	terminal.Read()
	return m
}

func (m *TmuxMenu) uninstall() Menu {
	m.Module.Uninstall()
	terminal.Read()
	return m
}

func (m *TmuxMenu) backHandler() Menu {
	return m.Parent
}
