package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Vim struct {
	printer *Printer
}

func (Vim) Name() string {
	return "Vim"
}

func (v Vim) Install(configuration Configuration) {
	v.printer.Print("Installing...")
	v.printer.Print("Adding .vimrc file to homedir")
	err := ioutil.WriteFile(v.filePath(), []byte(VimConf), 644)
	if err != nil {
		log.Fatalln("Error installing vim.", err)
	}
	v.printer.Print("Done!")
}

func (v Vim) Uninstall(configuration Configuration) {
	v.printer.Print("Uninstalling...")
	v.printer.Print("Removing .vimrc file to homedir")
	err := os.Remove(v.filePath())
	if err != nil {
		log.Fatalln("Error uninstalling vim.", err)
	}
	v.printer.Print("Done!")
}

func (v *Vim) SetPrinter(printer *Printer) {
	v.printer = printer
}

func (Vim) filePath() string {
	return fmt.Sprintf("%s/%s", homedir, VimConfFileName)
}
