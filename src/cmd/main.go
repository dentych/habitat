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

const HashKey = "#####"

const EnvConfFileName = ".env.conf"

var EnvConfFilePath string

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	EnvConfFilePath = fmt.Sprintf("%s/%s", internal.HomeDir, EnvConfFileName)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Subcommand is needed.")
		os.Exit(1)
	}

	conf := buildConfig()
	modules := buildModules(conf)

	switch os.Args[1] {
	case "install":
		installEnv(modules, conf)
	case "remove":
		removeEnv(modules, conf)
	case "reconfigure":
		fmt.Printf("Reconfiguring ENV\n\n")
		wizard := confwiz.New(&conf)
		conf := wizard.Configure()
		writeError := ioutil.WriteFile(EnvConfFilePath, conf.Marshal(), 0644)
		if writeError != nil {
			log.Fatalln("Couldn't save configuration file.", writeError)
		}
	case "removeconf":
		err := os.Remove(EnvConfFilePath)
		if err != nil {
			log.Fatalf("Error when removing ENV config file: %s", err)
		}
	default:
		fmt.Println("Invalid subcommand.")
		fmt.Println("Valid subcommands are: install, remove, reconfigure, removeconf")
		os.Exit(1)
	}
}

// This part is ugly
func buildConfig() internal.Configuration {
	fmt.Println("Looking for existing config file.")
	content, readError := ioutil.ReadFile(EnvConfFilePath)
	if readError == nil {
		fmt.Println("Found config file!")
		conf := internal.Configuration{}
		unmarshalError := conf.Unmarshal(content)
		if unmarshalError != nil {
			log.Fatalf("Error when reading %s. You might need to delete the file: %s", EnvConfFileName, unmarshalError)
		}

		return conf
	} else {
		if errors.Is(readError, os.ErrNotExist) {
			fmt.Println("Didn't file a config file. Starting configuration wizard!")
			wizard := confwiz.New(nil)
			conf := wizard.Configure()
			writeError := ioutil.WriteFile(EnvConfFilePath, conf.Marshal(), 0644)
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

func buildModules(configuration internal.Configuration) []internal.Module {
	modules := make([]internal.Module, 0, 10)
	if configuration.ModulesEnabled.Tmux {
		modules = append(modules, &internal.Tmux{})
	}
	if configuration.ModulesEnabled.Vim {
		modules = append(modules, &internal.Vim{})
	}
	if configuration.ModulesEnabled.Git {
		modules = append(modules, &internal.Git{})
	}
	if configuration.ModulesEnabled.Bash {
		modules = append(modules, &internal.Bash{})
	}

	return modules
}

func installEnv(modules []internal.Module, conf internal.Configuration) {
	printer := internal.Printer{}
	printer.Key = HashKey
	printer.Print("Starting installation of ENV")
	for _, v := range modules {
		printer.Key = v.Name()
		v.SetPrinter(&printer)
		v.Install(conf)
	}
	printer.Key = HashKey
	printer.Print("ENV installation has finished!")
}

func removeEnv(modules []internal.Module, conf internal.Configuration) {
	printer := internal.Printer{}
	printer.Key = HashKey
	printer.Print("Starting uninstallation of ENV")
	for _, v := range modules {
		printer.Key = v.Name()
		v.SetPrinter(&printer)
		v.Uninstall(conf)
	}
	printer.Key = HashKey
	printer.Print("ENV uninstallation has finished!")
}
