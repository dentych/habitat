package internal

import (
	"encoding/json"
	"log"
)

type Configuration struct {
	ModulesEnabled struct {
		Git  bool `json:"git"`
		Tmux bool `json:"tmux"`
		Vim  bool `json:"vim"`
	} `json:"modulesEnabled"`

	Git struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"git"`
}

func NewDefaultConfiguration() *Configuration {
	conf := Configuration{}
	conf.ModulesEnabled.Git = true
	conf.ModulesEnabled.Tmux = true
	conf.ModulesEnabled.Vim = true

	conf.Git.Name = "Default Name"
	conf.Git.Email = "default@example.com"

	return &conf
}

func (c *Configuration) Marshal() []byte {
	result, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalln("Marshalling error.", err)
	}

	return result
}

func (c *Configuration) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}
