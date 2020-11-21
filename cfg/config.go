package cfg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
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
		yamlContent := getFileContentFromHomeDir()
		err := yaml.Unmarshal(yamlContent, &config)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
	return config
}

func getFileContentFromHomeDir() []byte{
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot get User Home directory. Error: " + err.Error())
	}
	appConfigPath := usrHomeDir + "/.2Passwd/config.yaml"
	fileContent, fileReadErr := ioutil.ReadFile(appConfigPath)
	if fileReadErr != nil{
		_, ok := fileReadErr.(*os.PathError)
		if !ok {
			panic(fileReadErr)
		} else {
			//Check if folder exists
			_, err := ioutil.ReadDir(usrHomeDir + "/.2Passwd")
			if err != nil {
				_, ok := err.(*os.PathError)
				if !ok {
					panic(err)
				} else {
					//Create Path and directory
					err := os.Mkdir(usrHomeDir + "/.2Passwd", 0700)
					if err != nil {
						panic("Cannot create directory " + usrHomeDir + "/.2Passwd - Error: " + err.Error())
					}
					err = ioutil.WriteFile(appConfigPath, []byte("accounts:\n"), 0600)
					if err != nil {
						panic("Configuration file does not exist in " + appConfigPath + " and could not create it - Error: " + err.Error())
					}
				}
			} else {
				err = ioutil.WriteFile(appConfigPath, []byte("accounts:\n"), 0600)
				if err != nil {
					panic("Configuration file does not exist in " + appConfigPath + " and could not create it - Error: " + err.Error())
				}
			}
		}
		fileContent, fileReadErr = ioutil.ReadFile(appConfigPath)
		if fileReadErr != nil {
			panic("Could not read configuration file " + appConfigPath + " Error: " + fileReadErr.Error())
		}
	}
	return fileContent
}
