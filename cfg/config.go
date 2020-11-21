package cfg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var config *Config

type Config struct {
	Accounts []Account `yaml:"accounts"`
}
type Account struct {
	Name string `yaml:"name"`
	Seed string `yaml:"seed"`
	Star bool   `yaml:"star,omitempty"`
}

func Configuration() *Config{
	if config == nil {
		log.Printf("Loadgin accounts from configuration.yaml")
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}

		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}

	return config
}
