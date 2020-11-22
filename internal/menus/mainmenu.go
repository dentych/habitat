package menus

import (
	"fmt"
	"os"
)

type MainMenu struct {
	DefaultMenu
}

func NewMainMenu() *MainMenu {
	menu := MainMenu{}
	menu.Name = "Main menu"
	menu.Options = []Option{
		{Key: '1', Description: "Configuration", Handler: menu.configHandler},
		{Key: '2', Description: "Git module", Handler: menu.gitHandler},
		{Key: '3', Description: "Bash module", Handler: menu.bashHandler},
		{Key: '4', Description: "Tmux module", Handler: menu.tmuxHandler},
		{Key: 'q', Description: "Quit", Handler: menu.quitHandler},
	}

	return &menu
}

func (*MainMenu) quitHandler() Menu {
	fmt.Println("Good bye!")
	os.Exit(0)
	return nil
}

func (m *MainMenu) configHandler() Menu {
	return NewConfigMenu(m)
}

func (m *MainMenu) gitHandler() Menu {
	return NewGitMenu(m)
}

func (m *MainMenu) bashHandler() Menu {
	return NewBashMenu(m)
}

func (m *MainMenu) tmuxHandler() Menu {
	return NewTmuxMenu(m)
}
