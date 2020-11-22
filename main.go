package main

import (
	"gitlab.com/dentych/env/internal/menus"
)

var currentMenu menus.Menu

func main() {
	for {
		if currentMenu == nil {
			currentMenu = menus.NewMainMenu()
		}
		nextMenu := currentMenu.Execute()
		if nextMenu != nil {
			currentMenu = nextMenu
		}
	}
}
