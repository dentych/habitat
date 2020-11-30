package modules

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Tmux struct {
	homeDir string
}

func NewTmuxModule(homeDir string) Tmux {
	t := Tmux{}
	t.homeDir = homeDir
	return t
}

func (Tmux) Name() string {
	return "tmux"
}

func (t Tmux) Install() {
	fmt.Println("Installing...")
	fmt.Println("Adding .tmux.conf file to HomeDir")
	err := ioutil.WriteFile(t.filePath(), []byte(TmuxConf), 0644)
	if err != nil {
		log.Fatalln("Tmux installation failed.", err)
	}
	fmt.Println("Done!")
}

func (t Tmux) Uninstall() {
	fmt.Println("Uninstalling...")
	fmt.Println("Removing .tmux.conf file to HomeDir")
	err := os.Remove(t.filePath())
	if err != nil && !errors.Is(err, os.ErrNotExist){
		log.Fatalln("Tmux uninstallation error.", err)
	}
	fmt.Println("Done!")
}

func (t Tmux) filePath() string {
	return fmt.Sprintf("%s/%s", t.homeDir, TmuxConfFileName)
}
