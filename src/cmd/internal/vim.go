package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Vim struct {
	printer *Printer
}

func (Vim) Name() string {
	return "vim"
}

func (v Vim) Install(configuration Configuration) {
	v.printer.Print("Installing...")
	v.printer.Print("Adding .vimrc file to HomeDir")
	err := ioutil.WriteFile(v.filePath(), []byte(VimConf), 0644)
	if err != nil {
		log.Fatalln("Error installing vim.", err)
	}
	v.printer.Print("Done!")
}

func (v Vim) Uninstall(configuration Configuration) {
	v.printer.Print("Uninstalling...")
	v.printer.Print("Removing .vimrc file to HomeDir")
	err := os.Remove(v.filePath())
	if err != nil && !errors.Is(err, os.ErrNotExist){
		log.Fatalln("Error uninstalling vim.", err)
	}
	v.printer.Print("Done!")
}

func (v *Vim) SetPrinter(printer *Printer) {
	v.printer = printer
}

func (Vim) filePath() string {
	return fmt.Sprintf("%s/%s", HomeDir, VimConfFileName)
}
