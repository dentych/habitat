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

func addFileToBashrc(filename string) {
	bashrcPath := fmt.Sprintf("%s/.bashrc", HomeDir)
	strToAppend := fmt.Sprintf("\n. ~/%s\n", filename)
	content, err := ioutil.ReadFile(bashrcPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = ioutil.WriteFile(bashrcPath, []byte(strToAppend), 644)
			if err == nil {
				return
			}
		}
		log.Fatalln("Could not write bashrc file", err)
	} else {
		if !strings.Contains(string(content), filename) {
			f, fileErr := os.OpenFile(bashrcPath, os.O_APPEND | os.O_WRONLY, 644)
			if fileErr != nil {
				log.Fatalln("Could not write bashrc file [2]", fileErr)
			}
			defer f.Close()

			_, writeErr := f.WriteString(strToAppend)
			if writeErr != nil {
				log.Fatalf("Could not add %s to bashrc file. Error: %s", filename, err)
			}
		}
	}
}
