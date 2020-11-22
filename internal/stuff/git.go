package stuff

import (
	"errors"
	"gitlab.com/dentych/env/internal/configuration"
	"log"
	"os/exec"
)

var gitAliases = map[string]string{
	"alias.cp":    "cherry-pick",
	"alias.co":    "checkout",
	"alias.cl":    "clone",
	"alias.c":     "commit",
	"alias.st":    "status -sb",
	"alias.br":    "branch",
	"alias.d":     "diff",
	"alias.dc":    "diff --cached",
	"alias.p":     "pull -p",
	"alias.pu":    "push -u",
	"alias.f":     "fetch -p",
	"alias.b":     "branch",
	"alias.logn":  "log --all --graph --oneline --decorate",
	"alias.lognb": "log --graph --oneline --decorate",
	"alias.pushb": "!git push -u origin $(git rev-parse --abbrev-ref HEAD)",
}

type Git struct {
	printer *Printer
}

func (Git) Name() string {
	return "git"
}

func (g Git) Install(configuration configuration.Configuration) {
	g.printer.Print("Installing...")
	if !gitExists() {
		log.Fatalln("Git command not found. Please install git to use this module.")
	}

	g.printer.Print("Setting up username and email")
	executeCommand("user.name", configuration.Git.Name)
	executeCommand("user.email", configuration.Git.Email)

	g.printer.Print("Setting up git aliases")
	for k, v := range gitAliases {
		executeCommand(k, v)
	}

	g.printer.Print("Done!")
}

func (Git) Uninstall(configuration configuration.Configuration) {
	executeCommand("--unset", "user.name")
	executeCommand("--unset", "user.email")

	for k := range gitAliases {
		executeCommand("--unset", k)
	}
}

func (g *Git) SetPrinter(printer *Printer) {
	g.printer = printer
}

func gitExists() bool {
	cmd := exec.Command("git", "version")
	_, err := cmd.Output()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return false
		} else {
			log.Fatalln("Error executing command", err)
		}
	}
	return true
}

func executeCommand(args ...string) {
	fullArgs := []string{"config", "--global"}
	fullArgs = append(fullArgs, args...)
	cmd := exec.Command("git", fullArgs...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run command '%s': %s", cmd.String(), err)
	}
}
