package main

import (
	log "github.com/sirupsen/logrus"
	config "weatherapi/v2/cmd/svr/config"
	"weatherapi/v2/internal/facade"
	"weatherapi/v2/internal/mapper"
	"weatherapi/v2/internal/repository"
)

// InitUpstreamConfigClient initializes the upstream repository and injects it into the service.
func InitUpstreamConfigClient(cfg *config.Config) (*facade.Service, error) {
	// Retrieve the upstream service configuration
	upstreamConfig, err := cfg.GetService("WeatherService")
	if err != nil {
		log.Printf("Failed to establish the upstream service: %v", err)
		return nil, err
	}

	// Initialize the repository with the upstreamConfig
	repo := &repository.Repository{
		Config: upstreamConfig,
	}

	// Inject the repository and mapper into the facade.Service
	service := &facade.Service{
		RepositoryService: repo,
		Mapper:            &mapper.Mapper{},
	}

	return service, nil
}
