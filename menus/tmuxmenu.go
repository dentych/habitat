package menus

import (
	"gitlab.com/dentych/env/modules"
)

type TmuxMenu struct {
	DefaultMenu
}

func NewTmuxMenu(parent Menu, homeDir string) *TmuxMenu {
	menu := TmuxMenu{}
	menu.Name = "Tmux menu"
	menu.Parent = parent
	menu.Options = []Option{
		{Key: '1', Description: "Install", Handler: menu.Install},
		{Key: '2', Description: "Uninstall", Handler: menu.Uninstall},
		{Key: 'q', Description: "Back", Handler: menu.backHandler},
	}
	menu.Module = modules.NewTmuxModule(homeDir)

	return &menu
}

func (m *TmuxMenu) backHandler() Menu {
	return m.Parent
}
