package internal

import (
	"encoding/json"
	"log"
)

type Configuration struct {
	Git struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"git"`
}

func (c *Configuration) Marshal() []byte {
	result, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalln("Marshalling error.", err)
	}

	return result
}

func (c *Configuration) Unmarshal(data []byte) Configuration {
	var output Configuration
	err := json.Unmarshal(data, &output)
	if err != nil {
		log.Fatalln("Unmarshalling error.", err)
	}

	return output
}
