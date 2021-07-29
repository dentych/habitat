package configuration

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Configuration struct {
	Git struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Directory string `json:"directory"`
	} `json:"git"`
}

var Config Configuration

func (c *Configuration) Load() {
	data, err := ioutil.ReadFile(homeDir() + "/.habitat/config")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			c.Save()
		} else {
			log.Fatal("Failed to read configuration file: ", err)
		}
	} else {
		err = c.unmarshal(data)
		if err != nil {
			log.Fatal("Failed to load configuration file: ", err)
		}
	}
}

func (c *Configuration) Save() {
	err := ioutil.WriteFile(homeDir()+"/.habitat/config", Config.marshal(), 0644)
	if err != nil {
		log.Fatal("Failed to write configuration file: ", err)
	}
}

func (c *Configuration) marshal() []byte {
	result, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalln("Marshalling error.", err)
	}

	return result
}

func (c *Configuration) unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

func homeDir() string {
	homeDir, _ := os.UserHomeDir()
	return strings.Replace(homeDir, "\\", "/", -1)
}
