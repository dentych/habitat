package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var bashAliases = map[string]string{
	"gs":    "git status",
	"ls":    "ls --color",
	"vi":    "vim",
	"gc":    "git clean -f && git clean -f -d",
	"gca":   "git clean -f && git clean -f -d && git checkout -f",
	"dc":    "docker-compose",
	"drma":  "docker rm -f \\$(docker ps -aq)",
}

const bashFileName = ".env.bash-setup.sh"
const gitPromptFileName = ".env.git-prompt.sh"

type Bash struct {
	printer *Printer
}

func (Bash) Name() string {
	return "bash"
}

func (b Bash) Install(configuration Configuration) {
	b.printer.Print("Installing...")
	b.printer.Print("Generating bash aliases and PS1 config")
	var output bytes.Buffer

	// Special alias (that requires variable)
	bashAliases["cdgit"] = fmt.Sprintf("cd %s/Documents/git", HomeDir)

	for k, v := range bashAliases {
		output.WriteString(fmt.Sprintf("alias %s=\"%s\"\n", k, v))
	}

	output.WriteString(fmt.Sprintf("\nsource ~/%s\n", gitPromptFileName))
	output.WriteString("PS1='\\[\\e[32m\\]\\u@\\h \\[\\e[33m\\]\\w\\[\\e[92m\\]$(__git_ps1 \" (%s)\")\\[\\e[00m\\] $ '\n")

	b.printer.Print("Creating git prompt script file")
	err1 := ioutil.WriteFile(fmt.Sprintf("%s/%s", HomeDir, gitPromptFileName), []byte(GitPrompt), 0644)
	b.printer.Print("Creating bash setup script file")
	err2 := ioutil.WriteFile(fmt.Sprintf("%s/%s", HomeDir, bashFileName), output.Bytes(), 0644)
	if err1 != nil || err2 != nil {
		log.Fatalln("Could not write files:", err1, err2)
	}

	b.printer.Print("Adding bash setup script file to .bashrc file")
	addFileToBashrc(bashFileName)
	b.printer.Print("Done!")
}

func (Bash) Uninstall(configuration Configuration) {
	_ = os.Remove(fmt.Sprintf("%s/%s", HomeDir, gitPromptFileName))
	_ = os.Remove(fmt.Sprintf("%s/%s", HomeDir, bashFileName))
}

func (b *Bash) SetPrinter(printer *Printer) {
	b.printer = printer
}
