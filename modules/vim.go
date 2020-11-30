package modules

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Vim struct {
	homeDir string
}

func NewVimModule(homeDir string) Vim {
	v := Vim{}
	v.homeDir = homeDir
	return v
}

func (Vim) Name() string {
	return "vim"
}

func (v Vim) Install() {
	fmt.Println("Installing...")
	fmt.Println("Adding .vimrc file to HomeDir")
	err := ioutil.WriteFile(v.filePath(), []byte(VimConf), 0644)
	if err != nil {
		log.Fatalln("Error installing vim.", err)
	}
	fmt.Println("Done!")
}

func (v Vim) Uninstall() {
	fmt.Println("Uninstalling...")
	fmt.Println("Removing .vimrc file to HomeDir")
	err := os.Remove(v.filePath())
	if err != nil && !errors.Is(err, os.ErrNotExist){
		log.Fatalln("Error uninstalling vim.", err)
	}
	fmt.Println("Done!")
}

func (v Vim) filePath() string {
	return fmt.Sprintf("%s/%s", v.homeDir, VimConfFileName)
}
