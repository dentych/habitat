package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var HomeDir string

func init() {
	HomeDir, _ = os.UserHomeDir()
	HomeDir = strings.ReplaceAll(HomeDir, "\\", "/") // Fuck u Windows
}

func AddFileToBashrc(filepath string) {
	bashrcPath := fmt.Sprintf("%s/.bashrc", HomeDir)
	strToAppend := fmt.Sprintf("\n. %s\n", filepath)
	content, err := ioutil.ReadFile(bashrcPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = ioutil.WriteFile(bashrcPath, []byte(strToAppend), 0644)
			if err == nil {
				return
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
