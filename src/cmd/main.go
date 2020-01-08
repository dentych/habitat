package main

import (
	"env/cmd/internal"
	"env/cmd/internal/confwiz"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Subcommand is needed.")
		os.Exit(1)
	}

	modules := buildModules()

	switch os.Args[1] {
	case "install":
		installEnv(modules, buildConfig())
	case "remove":
		removeEnv(modules, buildConfig())
	case "reconfigure":
		fmt.Printf("Reconfiguring ENV\n\n")
		wizard := confwiz.New(nil)
		conf := wizard.Configure()
		writeError := ioutil.WriteFile(internal.HomeDir + "/.env.conf", conf.Marshal(), 644)
		if writeError != nil {
			log.Fatalln("Couldn't save configuration file.", writeError)
		}
	default:
		fmt.Println("Invalid subcommand.")
		os.Exit(1)
	}
}

func buildConfig() internal.Configuration {
	fmt.Println("Looking for existing config file.")
	content, readError := ioutil.ReadFile(internal.HomeDir + "/.env.conf")
	if readError == nil {
		fmt.Println("Found config file!")
		conf := internal.Configuration{}
		unmarshalError := conf.Unmarshal(content)
		if unmarshalError != nil {
			log.Fatalln("Error when reading .env.conf. You might need to delete the file.", unmarshalError)
		}

		return conf
	} else {
		if errors.Is(readError, os.ErrNotExist) {
			fmt.Println("Didn't file a config file. Starting configuration wizard!")
			wizard := confwiz.New(nil)
			conf := wizard.Configure()
			writeError := ioutil.WriteFile(internal.HomeDir + "/.env.conf", conf.Marshal(), 644)
			if writeError != nil {
				log.Fatalln("Couldn't save configuration file.", writeError)
			}
			return conf
		} else {
			log.Fatalln("Error when reading configuration file.", readError)
		}
	}
	return internal.Configuration{}
}

func buildModules() []internal.Module {
	modules := make([]internal.Module, 0, 10)
	modules = append(modules, &internal.Tmux{})
	modules = append(modules, &internal.Vim{})
	modules = append(modules, &internal.Git{})

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
