package menus

import (
	"gitlab.com/dentych/habitat/modules"
)

type BashMenu struct {
	DefaultMenu
}

func NewBashMenu(parent Menu, homeDir string) *BashMenu {
	menu := BashMenu{}
	menu.Name = "Bash menu"
	menu.Parent = parent
	menu.Module = modules.NewBashModule(homeDir)
	menu.Options = []Option{
		{Key: '1', Description: "Install", Handler: menu.Install},
		{Key: '2', Description: "Uninstall", Handler: menu.Uninstall},
		{Key: 'q', Description: "Back", Handler: menu.Back},
	}

	return &menu
}
