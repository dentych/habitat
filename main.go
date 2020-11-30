package main

import (
	"errors"
	"gitlab.com/dentych/env/configuration"
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
	ensureEnvFolderExists(homeDir)
	configuration.Config = configuration.Configuration{}
	configuration.Config.Load()
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

func ensureEnvFolderExists(homeDir string) {
	_, err := os.Open(homeDir + "/.env")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatal("Failed to open .env folder", err)
		}
		err := os.Mkdir(homeDir +  "/.env", 0755)
		if err != nil {
			log.Fatal("Failed to create .env directory", err)
		}
	}
}
