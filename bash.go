package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gitlab.com/dentych/habitat/terminal"
)

var bashAliases = map[string]string{
	"gs":   "git status",
	"ls":   "ls -G",
	"vi":   "vim",
	"gc":   "git clean -f && git clean -f -d",
	"gca":  "git clean -f && git clean -f -d && git checkout -f",
	"dc":   "docker-compose",
	"drma": "docker rm -f \\$(docker ps -aq)",
	"got":  "go test ./...",
	"goti": "go test --tags integration ./...",
}

const bashFileName = "bash-setup.sh"
const gitPromptFileName = "git-prompt.sh"
const customBashFileName = "custom.sh"

const bashSourceString = "\n. %s\n"

func installBash(homeDir string) {
	terminal.PrintHeading("Bash")
	fmt.Println("Installing...")
	fmt.Println("Generating bash aliases and PS1 config")
	var output bytes.Buffer

	resp, err := http.Get("https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh")
	if err != nil {
		log.Fatalf("Failed to download git-prompt: %s", err)
	}
	defer resp.Body.Close()

	fileInfo, err := os.Stat(fmt.Sprintf("%s/.habitat", homeDir))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.Mkdir(fmt.Sprintf("%s/.habitat", homeDir), 0755)
			if err != nil {
				log.Fatalf("Failed to create .habitat directory")
			}
		} else {
			log.Fatalf("Failed to check stats on .habitat directory")
		}
	}

	if fileInfo != nil && !fileInfo.IsDir() {
		log.Fatalf(".habitat exists but is not a directory")
	}

	out, err := os.Create(fmt.Sprintf("%s/.habitat/%s", homeDir, gitPromptFileName))
	if err != nil {
		log.Fatalf("Failed to write git-prompt.sh file: %s", err)
	}
	defer out.Close()

	customFilePath := fmt.Sprintf("%s/.habitat/%s", homeDir, customBashFileName)
	_, err = os.Stat(customFilePath)
	if os.IsNotExist(err) {
		f, err := os.Create(customFilePath)
		if err != nil {
			log.Fatalf("Failed to write custom bash setup file: %s", err)
		}
		defer f.Close()
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Failed to write to git-prompt file: %s", err)
	}

	// Special alias (that requires variable)
	bashAliases["cdgit"] = fmt.Sprintf("cd %s/Documents/git", homeDir)

	for k, v := range bashAliases {
		output.WriteString(fmt.Sprintf("alias %s=\"%s\"\n", k, v))
	}

	output.WriteString(fmt.Sprintf("\nsource %s/.habitat/%s\n", homeDir, gitPromptFileName))
	output.WriteString("PS1='\\[\\e[32m\\]\\u@\\h \\[\\e[33m\\]\\w\\[\\e[92m\\]$(__git_ps1 \" (%s)\")\\[\\e[00m\\] $ '\n")
	output.WriteString(fmt.Sprintf("\n. %s/.habitat/%s\n", homeDir, customBashFileName))

	fmt.Println("Creating bash setup script file")
	err = ioutil.WriteFile(fmt.Sprintf("%s/.habitat/%s", homeDir, bashFileName), output.Bytes(), 0644)
	if err != nil {
		log.Fatalln("Could not write bash-setup.sh file:", err)
	}

	fmt.Println("Adding bash setup script file to .bashrc file")
	addFileToBash(homeDir, fmt.Sprintf("%s/.habitat/%s", homeDir, bashFileName))
	fmt.Println("Done!")
}

func addFileToBash(homeDir string, filepath string) {
	bashrcPath := fmt.Sprintf("%s/.zshrc", homeDir)
	_, err := os.Stat(bashrcPath)
	if err != nil {
		bashrcPath = fmt.Sprintf("%s/.bashrc", homeDir)
	}
	strToAppend := fmt.Sprintf(bashSourceString, filepath)
	content, err := ioutil.ReadFile(bashrcPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Did not find .bashrc in home directory. Creating now")
			err = ioutil.WriteFile(bashrcPath, []byte(strToAppend), 0644)
			if err == nil {
				return
			} else {
				log.Fatalln("Failed to write bash file", err)
			}
		} else {
			log.Fatalln("Could not write bashrc file", err)
		}
	} else {
		if !strings.Contains(string(content), filepath) {
			f, fileErr := os.OpenFile(bashrcPath, os.O_APPEND|os.O_WRONLY, 0644)
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
