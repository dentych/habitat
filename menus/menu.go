package menus

import (
	"fmt"
	"gitlab.com/dentych/habitat/modules"
	"gitlab.com/dentych/habitat/terminal"
)

var ErrEmptyInput = fmt.Errorf("empty input")

type Menu interface {
	Execute() Menu
}

type DefaultMenu struct {
	Name    string
	Parent  Menu
	Module  modules.Module
	Options []Option
}

func (m *DefaultMenu) FindOption(key byte) *Option {
	for _, v := range m.Options {
		if v.Key == key {
			return &v
		}
	}
	return nil
}

func (m *DefaultMenu) Install() Menu {
	m.Module.Install()
	terminal.ReadEnter()
	return m.Parent
}

func (m *DefaultMenu) Uninstall() Menu {
	m.Module.Uninstall()
	terminal.ReadEnter()
	return m.Parent
}

func (m *DefaultMenu) Back() Menu {
	return m.Parent
}

func (m *DefaultMenu) Execute() Menu {
	terminal.Clear()
	fmt.Println("---- " + m.Name + " ----")
	m.PrintOptions()
	input := terminal.ReadByte()
	option := m.FindOption(input)
	if option != nil {
		return option.Handle()
	}
	return nil
}

func (m *DefaultMenu) PrintOptions() {
	for _, v := range m.Options {
		fmt.Printf("%c. %v\n", v.Key, v.Description)
	}
	fmt.Println()
}
