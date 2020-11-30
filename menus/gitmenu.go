package menus

import (
	"gitlab.com/dentych/env/modules"
	"gitlab.com/dentych/env/terminal"
)

type GitMenu struct {
	DefaultMenu
	GitModule modules.Module
}

func NewGitMenu(parent Menu) *GitMenu {
	menu := GitMenu{}
	menu.Name = "git menu"
	menu.Parent = parent
	menu.Options = []Option{
		{Key: '1', Description: "Install", Handler: menu.install},
		{Key: '2', Description: "Uninstall", Handler: menu.uninstall},
		{Key: 'q', Description: "Back", Handler: func() Menu { return menu.Parent }},
	}
	menu.GitModule = modules.NewGitModule()

	return &menu
}

func (m *GitMenu) install() Menu {
	m.GitModule.Install()
	terminal.Read()
	return m
}

func (m *GitMenu) uninstall() Menu {
	m.GitModule.Uninstall()
	terminal.Read()
	return m
}
