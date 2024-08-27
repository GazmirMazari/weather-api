package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"strings"
	"time"
)

type Config struct {
	Env            string          `yaml:"Env"`
	Port           string          `yaml:"Port"`
	AppName        string          `yaml:"AppName"`
	ClientConfig   ClientConfig    `yaml:"ClientConfig"`
	ServiceConfigs []ServiceConfig `yaml:"ServiceConfigs"` // Updated to match array in YAML
}

type ClientConfig struct {
	Timeout            int `yaml:"Timeout"`
	IdleConnTimeout    int `yaml:"IdleConnTimeout"`
	MaxIdleConsPerHost int `yaml:"MaxIdleConsPerHost"`
	MaxConsPerHost     int `yaml:"MaxConsPerHost"`
}

type ServiceConfig struct {
	Name                   string                 `yaml:"Name"`
	URL                    string                 `yaml:"URL"`
	Timeout                int                    `yaml:"Timeout"`
	mergedComponentConfigs map[string]interface{} // Not from YAML
	Client                 *http.Client           // Not from YAML
}

// Method to initialize the HTTP client for each service
func (s *ServiceConfig) InitHTTPClient() {
	s.Client = &http.Client{
		Timeout: time.Duration(s.Timeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:    100,
			IdleConnTimeout: 90 * time.Second,
		},
	}
}

type ServiceConfigMap map[string]*ServiceConfig

// Load initializes and returns a configuration object.
func New(configPath string) (*Config, error) {
	log.Tracef("Loading config from: %s\n", configPath)

	// Load the config from the YAML file
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

	// Initialize HTTP clients for all services
	for name, service := range config.ServiceConfigs {
		log.Tracef("Initializing service: %s", name)
		service.InitHTTPClient()
	}

	log.Tracef("Environment: %s\n", strings.ToUpper(config.Env))
	return config, nil
}

func (c *Config) GetService(name string) (*ServiceConfig, error) {
	for _, service := range c.ServiceConfigs {
		if service.Name == name {
			log.Tracef("Service %q found: %+v", name, service)
			return &service, nil
		}
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
