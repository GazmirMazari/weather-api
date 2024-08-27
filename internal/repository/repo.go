package repository

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
	"weatherapi/v2/cmd/svr/config"
	"weatherapi/v2/external/models"
)

type RepositoryI interface {
	SearchWeatherApi(ctx context.Context, request models.Request) (response models.WeatherPropertiesRes, err error)
}

// Repository struct that uses the ServiceConfig and implements RepositoryI
type Repository struct {
	Config *config.ServiceConfig // Injected ServiceConfig
}

// Ensure Repository implements RepositoryI
func (r *Repository) SearchWeatherApi(ctx context.Context, request models.Request) (models.WeatherPropertiesRes, error) {
	// Build the API URL using the ServiceConfig
	apiURL := fmt.Sprintf("%s/%s,%s", r.Config.URL, request.Latitude, request.Longitude)
	// Set the timeout based on the configuration
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.Config.Timeout)*time.Second)
	defer cancel() // Ensure the context is canceled after the request is done
	// Make the HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return models.WeatherPropertiesRes{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("failed to call the upstream service %v", err)
		return models.WeatherPropertiesRes{}, err
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("failed to read response body %v", err)
		return models.WeatherPropertiesRes{}, err
	}

	var weatherRes models.WeatherPropertiesRes
	if err := json.Unmarshal(body, &weatherRes); err != nil {
		log.Errorf("failed to unmarshal response body %v", err)
		return models.WeatherPropertiesRes{}, err
	}

	return weatherRes, nil
}
