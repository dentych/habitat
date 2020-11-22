package menus

import "gitlab.com/dentych/env/internal/modules"

type TmuxMenu struct {
	DefaultMenu
	TmuxModule modules.Module
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
}

func (m *TmuxMenu) install() Menu {
	return m
}

func (m *TmuxMenu) uninstall() Menu {
	return m
}

func (m *TmuxMenu) backHandler() Menu {
	return m.Parent
}
