package menus

import (
	"fmt"
	"gitlab.com/dentych/env/internal/terminal"
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

func (m *MainMenu) print() {
	terminal.Clear()
	m.PrintOptions()
}

func (m *MainMenu) gitHandler() Menu {
	// TODO
	panic("OMG")
}
