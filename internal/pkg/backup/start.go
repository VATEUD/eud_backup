package backup

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// ReadConfigFile reads the config file and parses it to the Config struct which holds all the databases list
func ReadConfigFile() (Config, error) {
	bytes, err := ioutil.ReadFile("configs/config.yaml")

	if err != nil {
		return Config{}, err
	}

	var config Config

	if err = yaml.Unmarshal(bytes, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
