package internal

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

type Git struct {
	printer *Printer
}

func (Git) Name() string {
	return "git"
}

func (g Git) Install(configuration Configuration) {

	cmd := exec.Command("git", "version")

	output, err := cmd.Output()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			g.printer.Print("ERROR: git command not found! Please install git to use this.")
			return
		} else {
			log.Fatalln("Error installing git.", err)
		}
	}

	result := strings.TrimSuffix(string(output), "\n")
	g.printer.Print(result)
}

func (Git) Uninstall(configuration Configuration) {

}

func (g *Git) SetPrinter(printer *Printer) {
	g.printer = printer
}
