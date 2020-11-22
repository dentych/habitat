package modules

import (
	"bytes"
	"fmt"
	"gitlab.com/dentych/env/internal"
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

const bashFileName = "bash-setup.sh"
const gitPromptFileName = "git-prompt.sh"

type bash struct {
}

func NewBashModule() Module {
	return bash{}
}

func (bash) Name() string {
	return "bash"
}

func (b bash) Install() {
	fmt.Println("Installing...")
	fmt.Println("Generating bash aliases and PS1 config")
	var output bytes.Buffer

	// Special alias (that requires variable)
	bashAliases["cdgit"] = fmt.Sprintf("cd %s/Documents/git", internal.HomeDir)

	for k, v := range bashAliases {
		output.WriteString(fmt.Sprintf("alias %s=\"%s\"\n", k, v))
	}

	output.WriteString(fmt.Sprintf("\nsource %s/.env/%s\n", internal.HomeDir, gitPromptFileName))
	output.WriteString("PS1='\\[\\e[32m\\]\\u@\\h \\[\\e[33m\\]\\w\\[\\e[92m\\]$(__git_ps1 \" (%s)\")\\[\\e[00m\\] $ '\n")

	fmt.Println("Creating git prompt script file")
	err1 := ioutil.WriteFile(fmt.Sprintf("%s/.env/%s", internal.HomeDir, gitPromptFileName), []byte(internal.GitPrompt), 0644)
	fmt.Println("Creating bash setup script file")
	err2 := ioutil.WriteFile(fmt.Sprintf("%s/.env/%s", internal.HomeDir, bashFileName), output.Bytes(), 0644)
	if err1 != nil || err2 != nil {
		log.Fatalln("Could not write files:", err1, err2)
	}

	fmt.Println("Adding bash setup script file to .bashrc file")
	internal.AddFileToBashrc(fmt.Sprintf("%s/.env/%s", internal.HomeDir, bashFileName))
	fmt.Println("Done!")
}

func (bash) Uninstall() {
	_ = os.Remove(fmt.Sprintf("%s/%s", internal.HomeDir, gitPromptFileName))
	_ = os.Remove(fmt.Sprintf("%s/%s", internal.HomeDir, bashFileName))
}
