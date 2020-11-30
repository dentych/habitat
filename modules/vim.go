package modules

import (
	"errors"
	"fmt"
	"gitlab.com/dentych/env"
	"gitlab.com/dentych/env/configuration"
	"io/ioutil"
	"log"
	"os"
)

type Vim struct {
}

func (Vim) Name() string {
	return "vim"
}

func (v Vim) Install(configuration configuration.Configuration) {
	fmt.Println("Installing...")
	fmt.Println("Adding .vimrc file to HomeDir")
	err := ioutil.WriteFile(v.filePath(), []byte(main.VimConf), 0644)
	if err != nil {
		log.Fatalln("Error installing vim.", err)
	}
	fmt.Println("Done!")
}

func (v Vim) Uninstall(configuration configuration.Configuration) {
	fmt.Println("Uninstalling...")
	fmt.Println("Removing .vimrc file to HomeDir")
	err := os.Remove(v.filePath())
	if err != nil && !errors.Is(err, os.ErrNotExist){
		log.Fatalln("Error uninstalling vim.", err)
	}
	fmt.Println("Done!")
}

func (Vim) filePath() string {
	return fmt.Sprintf("%s/%s", main.HomeDir, main.VimConfFileName)
}