package main

import (
	"env/cmd/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Subcommand is needed.")
		os.Exit(1)
	}

	modules := buildModules()

	conf := internal.Configuration{}

	switch os.Args[1] {
	case "install":
		installEnv(modules, conf)
	case "remove":
		removeEnv(modules, conf)
	default:
		fmt.Println("Invalid subcommand.")
		os.Exit(1)
	}
}

func buildModules() []internal.Module {
	modules := make([]internal.Module, 0, 10)
	modules = append(modules, &internal.Tmux{})
	modules = append(modules, &internal.Vim{})

	return modules
}

func installEnv(modules []internal.Module, conf internal.Configuration) {
	printer := internal.Printer{}
	printer.Key = "###"
	printer.Print("Starting installation of ENV")
	for _, v := range modules {
		printer.Key = v.Name()
		v.SetPrinter(&printer)
		v.Install(conf)
	}
	printer.Key = "###"
	printer.Print("ENV installation has finished!")
}

func removeEnv(modules []internal.Module, conf internal.Configuration) {
	fmt.Println("Uninstalling env...")
	for _, v := range modules {
		v.Uninstall(conf)
	}
	fmt.Println("Uninstallion completed!")
}

func printMethod(msg string) {
	fmt.Printf("- %s\n", msg)
}
