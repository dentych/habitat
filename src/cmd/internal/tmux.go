package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Tmux struct {
	printer *Printer
}

func (Tmux) Name() string {
	return "tmux"
}

func (t Tmux) Install(conf Configuration) {
	t.printer.Print("Installing...")
	t.printer.Print("Adding .tmux.conf file to HomeDir")
	err := ioutil.WriteFile(t.filePath(), []byte(TmuxConf), 644)
	if err != nil {
		log.Fatalln("Tmux installation failed.", err)
	}
	t.printer.Print("Done!")
}

func (t Tmux) Uninstall(conf Configuration) {
	t.printer.Print("Uninstalling...")
	t.printer.Print("Removing .tmux.conf file to HomeDir")
	err := os.Remove(t.filePath())
	if err != nil && !errors.Is(err, os.ErrNotExist){
		log.Fatalln("Tmux uninstallation error.", err)
	}
	t.printer.Print("Done!")
}

func (t *Tmux) SetPrinter(printer *Printer) {
	t.printer = printer
}

func (Tmux) filePath() string {
	return fmt.Sprintf("%s/%s", HomeDir, TmuxConfFileName)
}
