package repository

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"weatherapi/v2/cmd/svr/config"
	"weatherapi/v2/external/models"
)

type RepositoryI interface {
	SearchWeatherApi(ctx context.Context, request models.Request) (response models.WeatherPropertiesRes, err error)
}

type Repository struct {
	Config *config.Config
}

func (r *Repository) SearchWeatherApi(ctx context.Context, request models.Request) (response models.WeatherPropertiesRes, err error) {
	// Build the request URL
	requestUrl := fmt.Sprintf(r.Config.Services["NationalWeatherService"].URL+"%d"+"%d", request.Latitude, request.Longitude)

	// Create an HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", requestUrl, nil)
	if err != nil {
		log.WithError(err).Error("failed to create request %v", err)
		return response, err
	}

	// Execute the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("failed to execute request")
		return response, err
	}
	defer res.Body.Close()

	// Check for non-200 status codes
	if res.StatusCode != http.StatusOK {
		log.WithField("status_code", res.StatusCode).Error("unexpected status code")
		return response, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}

	// Unmarshal the response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.WithError(err).Error("failed to unmarshal the response %v", err)
		return response, err
	}

	return response, nil
}
