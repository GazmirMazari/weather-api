package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Config struct {
	AppName string `yaml:"AppName"`
	Env     string `yaml:"Env"`
	Port    string `yaml:"Port"`

	ComponentConfigs ComponentConfigs `yaml:"ComponentConfigs"`
	Services         ServiceConfigMap `yaml:"Services"`
}

type ComponentConfigs struct {
	//TODO add logging
	Client ClientConfig
}

func New(configPath string) (config *Config) {
	log.Tracef("config: %s\n", configPath)
	var errs []error
	if config, errs = new(builder).newConfig(configPath); len(errs) > 0 || config == nil {
		for _, err := range errs {
			log.Panicf("configuration error: %v\n", err.Error())
		}
		if config == nil {
			log.Panicln("configuration file not found")
		}
		log.Panicln("Exiting: failed to load the config file")
	}
	log.Tracef("env: %s\n", strings.ToUpper(config.Env))
	return config
}

// Service returns an initialized service configuration by name
func (c *Config) Service(name string) (*ServiceConfig, error) {
	if service, ok := c.Services[name]; ok {
		return service, nil
	}
	// return error if the service not found in config
	return nil, fmt.Errorf("Service: %s", fmt.Sprintf("%s not found", name))
}

func appendAndLog(err error, errs []error) []error {
	log.Error(err)
	return append(errs, err)
}
