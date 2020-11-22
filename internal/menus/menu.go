package menus

import (
	"fmt"
	"gitlab.com/dentych/env/internal/terminal"
)

var ErrEmptyInput = fmt.Errorf("empty input")

type Menu interface {
	Execute() Menu
}

type DefaultMenu struct {
	Name    string
	Parent  Menu
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

func (m *DefaultMenu) Execute() Menu {
	terminal.Clear()
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