package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	_ "gopkg.in/yaml.v2"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	AppName          string           `yaml:"AppName"`
	Env              string           `yaml:"Env"`
	Port             string           `yaml:"Port"`
	ComponentConfigs ComponentConfigs `yaml:"ComponentConfigs"`
	Services         ServiceConfigMap `yaml:"Services"`
}

type ComponentConfigs struct {
	Client ClientConfig `yaml:"Client"`
}

type ClientConfig struct {
	Timeout        int    `yaml:"Timeout"`
	BaseURL        string `yaml:"BaseURL"`
	MaxConnections int    `yaml:"MaxConnections"`
	RetryPolicy    string `yaml:"RetryPolicy"`
}

type ServiceConfig struct {
	URL                    string                 `yaml:"url"`
	Timeout                int                    `yaml:"timeout"`
	mergedComponentConfigs map[string]interface{} // Add this field
	Client                 *http.Client           // Add this field
	Name                   string                 // Add this field if needed
}

type ServiceConfigMap map[string]*ServiceConfig

// New initializes and returns a configuration object.
func New(configPath string) (*Config, error) {
	log.Tracef("Loading config from: %s\n", configPath)

	config, errs := newConfig(configPath)
	if len(errs) > 0 || config == nil {
		for _, err := range errs {
			log.Errorf("Configuration error: %v\n", err)
		}
		if config == nil {
			log.Errorf("Configuration file not found. Exiting.")
			return nil, fmt.Errorf("configuration file not found")
		}
		log.Errorf("Failed to load the config file. Exiting.")
		return nil, fmt.Errorf("failed to load the config file")
	}

	log.Tracef("Environment: %s\n", strings.ToUpper(config.Env))
	return config, nil
}

func (c *Config) GetService(name string) (*ServiceConfig, error) {
	if service, ok := c.Services[name]; ok {
		log.Tracef("Service %q found: %+v", name, service)
		return service, nil
	}
	// return error if the service not found in config
	return nil, fmt.Errorf("service %q not found", name)
}

// newConfig loads the configuration from the specified path.
func newConfig(configPath string) (*Config, []error) {
	var config Config
	var errs []error

	file, err := os.Open(configPath)
	if err != nil {
		return nil, appendAndLog(err, errs)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, appendAndLog(err, errs)
	}

	return &config, errs
}

// appendAndLog logs the error and appends it to the error list.
func appendAndLog(err error, errs []error) []error {
	log.Error(err)
	return append(errs, err)
}
