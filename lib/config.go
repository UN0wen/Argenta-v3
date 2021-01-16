package lib

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration is the base yaml object.
type Configuration struct {
	Database struct {
		URL      string `yaml:"url"`
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Bot struct {
		Token string `yaml:"token"`
		Owner string `yaml:"owner"`
		ID    string `yaml:"id"`
	} `yaml:"bot"`
	Prefix string `yaml:"prefix"`
}

// Config retrieves the app's configuration form config.json.
func Config() Configuration {
	file, err := os.Open("../config.yaml")
	CheckErr(err)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	CheckErr(err)

	var config Configuration
	err = yaml.Unmarshal(contents, &config)
	CheckErr(err)

	return config
}
