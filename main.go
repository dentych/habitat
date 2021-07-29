package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "embed"

	"gitlab.com/dentych/habitat/vscode"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "sync":
			sync()
		case "scrape":
			scrape()
		}
	}
}

func sync() {
	fmt.Println("Syncing...")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get home directory of user:", err)
	}

	fmt.Printf("Home directory: %s\n", homeDir)

	installBash(homeDir)
	installTmux(homeDir)
	installVim(homeDir)
	installGit()
	installVscode(homeDir)
}

func scrape() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get home directory of user:", err)
	}
	vscode.ScrapeVSCodeConfig(homeDir)
}

//go:embed tmux.conf
var tmuxConf []byte

const TmuxConfFilename = ".tmux.conf"

func installTmux(homeDir string) {
	fmt.Println("---------- Installing Tmux ----------")
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
	fmt.Println("---------- Installing Vim ----------")
	fmt.Println("Installing...")
	fmt.Println("Adding .vimrc file to HomeDir")
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s", homeDir, vimConfFilename), vimConf, 0644)
	if err != nil {
		log.Fatalln("Error installing vim.", err)
	}
	fmt.Println("Done!")

}

func installVscode(homeDir string) {
	fmt.Println("Installing VSCode configuration")
}
