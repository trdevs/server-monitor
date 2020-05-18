package configuration

import (
	"errors"
	"os"

	l "github.com/ermanimer/logger"
	"gopkg.in/yaml.v3"
)

const (
	ConfigurationFile = "./configuration/configuration.yaml"
)

func Load() (*Configuration, error) {
	cf, err := os.Open(ConfigurationFile)
	if err != nil {
		l.Debug(err.Error())
		return nil, errors.New("Opening configuration file failed!")
	}
	defer cf.Close()
	var c Configuration
	err = yaml.NewDecoder(cf).Decode(&c)
	if err != nil {
		l.Debug(err.Error())
		return nil, errors.New("Reading configuration file failed!")
	}
	return &c, nil
}
