package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "embed"

	"gitlab.com/dentych/habitat/terminal"
	"gitlab.com/dentych/habitat/vscode"
)

func main() {
	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "sync":
			sync()
		case "scrape":
			scrape()
		default:
			fmt.Println("Valid commands are:\n* sync - will sync all configuration files\n* scape - will scrape config files and save to the repo")
		}
	}
}

func sync() {
	terminal.PrintHeading("Synchronizing configuration")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Could not get home directory of user:", err)
	}

	fmt.Printf("Home directory is: %s\n", homeDir)

	installBash(homeDir)
	installTmux(homeDir)
	installVim(homeDir)
	installGit()
	installVscode(homeDir)
}

func scrape() {
	terminal.PrintHeading("Scraping configuration files")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Could not get home directory of user:", err)
	}
	vscode.ScrapeVSCode(homeDir)
}

//go:embed tmux.conf
var tmuxConf []byte

const TmuxConfFilename = ".tmux.conf"

func installTmux(homeDir string) {
	terminal.PrintHeading("Installing Tmux")
	fmt.Println("Adding .tmux.conf file to HomeDir")
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s", homeDir, TmuxConfFilename), []byte(tmuxConf), 0644)
	if err != nil {
		log.Fatalln("Tmux installation failed.", err)
	}
	fmt.Println("Done!")
}

//go:embed vim.conf
var vimConf []byte

var vimConfFilename = ".vimrc"

func installVim(homeDir string) {
	terminal.PrintHeading("Installing Vim")
	fmt.Println("Installing...")
	fmt.Println("Adding .vimrc file to HomeDir")
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s", homeDir, vimConfFilename), vimConf, 0644)
	if err != nil {
		log.Fatalln("Error installing vim.", err)
	}
	fmt.Println("Done!")

}

func installVscode(homeDir string) {
	terminal.PrintHeading("VSCode")
	fmt.Println("Installing VSCode configuration")
	vscode.InstallVSCode(homeDir)
}
