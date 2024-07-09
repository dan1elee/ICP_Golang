package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Database string `yaml:"database"`
}

var configuration *Configuration

func InitConfiguration(addr string) error {
	data, err := os.ReadFile(addr)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return err
}

func GetConfiguration() *Configuration {
	return configuration
}
