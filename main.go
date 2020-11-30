package main

import (
	"gitlab.com/dentych/env/menus"
	"log"
	"os"
	"strings"
)

var currentMenu menus.Menu

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get home directory of user", err)
	}
	homeDir = strings.Replace(homeDir, "\\", "/", -1)
	for {
		if currentMenu == nil {
			currentMenu = menus.NewMainMenu(homeDir)
		}
		nextMenu := currentMenu.Execute()
		if nextMenu != nil {
			currentMenu = nextMenu
		}
	}
}
