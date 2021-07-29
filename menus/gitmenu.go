package menus

import (
	"gitlab.com/dentych/habitat/modules"
)

type GitMenu struct {
	DefaultMenu
}

func NewGitMenu(parent Menu) *GitMenu {
	menu := GitMenu{}
	menu.Name = "git menu"
	menu.Parent = parent
	menu.Options = []Option{
		{Key: '1', Description: "Install", Handler: menu.Install},
		{Key: '2', Description: "Uninstall", Handler: menu.Uninstall},
		{Key: 'q', Description: "Back", Handler: menu.Back},
	}
	menu.Module = modules.NewGitModule()

	return &menu
}
