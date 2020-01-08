package internal

import "fmt"

type Module interface {
	Name() string

	// Installation script for the module
	Install(configuration Configuration)

	// Uninstallation script for the module
	Uninstall(configuration Configuration)

	SetPrinter(printer *Printer)
}

type Printer struct {
	Key string
}

func (p *Printer) Print(msg string) {
	fmt.Printf("[%s] %s\n", p.Key, msg)
}
