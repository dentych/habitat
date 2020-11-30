package modules

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

const bashSourceString = "\n. %s\n"

type bash struct {
	homeDir string
}

func NewBashModule(homeDir string) Module {
	module := bash{}
	module.homeDir = homeDir
	return module
}

func (bash) Name() string {
	return "bash"
}

func (b bash) Install() {
	fmt.Println("Installing...")
	fmt.Println("Generating bash aliases and PS1 config")
	var output bytes.Buffer

	// Special alias (that requires variable)
	bashAliases["cdgit"] = fmt.Sprintf("cd %s/Documents/git", b.homeDir)

	for k, v := range bashAliases {
		output.WriteString(fmt.Sprintf("alias %s=\"%s\"\n", k, v))
	}

	output.WriteString(fmt.Sprintf("\nsource %s/.env/%s\n", b.homeDir, gitPromptFileName))
	output.WriteString("PS1='\\[\\e[32m\\]\\u@\\h \\[\\e[33m\\]\\w\\[\\e[92m\\]$(__git_ps1 \" (%s)\")\\[\\e[00m\\] $ '\n")

	fmt.Println("Creating git prompt script file")
	err1 := ioutil.WriteFile(fmt.Sprintf("%s/.env/%s", b.homeDir, gitPromptFileName), []byte(GitPrompt), 0644)
	fmt.Println("Creating bash setup script file")
	err2 := ioutil.WriteFile(fmt.Sprintf("%s/.env/%s", b.homeDir, bashFileName), output.Bytes(), 0644)
	if err1 != nil || err2 != nil {
		log.Fatalln("Could not write files:", err1, err2)
	}

	fmt.Println("Adding bash setup script file to .bashrc file")
	b.addFileToBash(fmt.Sprintf("%s/.env/%s", b.homeDir, bashFileName))
	fmt.Println("Done!")
}

func (b bash) Uninstall() {
	fmt.Println("Removing git prompt file")
	_ = os.Remove(fmt.Sprintf("%s/.env/%s", b.homeDir, gitPromptFileName))
	fmt.Println("Removing bash setup file")
	_ = os.Remove(fmt.Sprintf("%s/.env/%s", b.homeDir, bashFileName))
	fmt.Println("Removing bash sourcing from .bashrc file")
	b.removeFileFromBash(fmt.Sprintf("%s/.env/%s", b.homeDir, bashFileName))
}

func (b bash) addFileToBash(filepath string) {
	bashrcPath := fmt.Sprintf("%s/.bashrc", b.homeDir)
	strToAppend := fmt.Sprintf(bashSourceString, filepath)
	content, err := ioutil.ReadFile(bashrcPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = ioutil.WriteFile(bashrcPath, []byte(strToAppend), 0644)
			if err == nil {
				return
			} else {
				log.Fatal("Failed to write bash file", err)
			}
		}
		log.Fatalln("Could not write bashrc file", err)
	} else {
		if !strings.Contains(string(content), filepath) {
			f, fileErr := os.OpenFile(bashrcPath, os.O_APPEND | os.O_WRONLY, 0644)
			if fileErr != nil {
				log.Fatalln("Could not write bashrc file [2]", fileErr)
			}
			defer f.Close()

			_, writeErr := f.WriteString(strToAppend)
			if writeErr != nil {
				log.Fatalf("Could not add %s to bashrc file. Error: %s", filepath, err)
			}
		}
	}
}

func (b bash) removeFileFromBash(filepath string) {
	bashrcPath := fmt.Sprintf("%s/.bashrc", b.homeDir)
	strToRemove := fmt.Sprintf(bashSourceString, filepath)
	content, err := ioutil.ReadFile(bashrcPath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatal("Could not read bash file", err)
		}
	}
	content = bytes.Replace(content, []byte(strToRemove), []byte(""), -1)
	err = ioutil.WriteFile(bashrcPath, content, 0644)
	if err != nil {
		log.Fatal("Failed to write .bashrc file", err)
	}
}