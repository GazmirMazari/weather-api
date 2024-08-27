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
	SearchWeatherApi(ctx context.Context, request models.Request) (WeatherResponse, error)
	GetGridInfo(ctx context.Context, request models.Request) (string, error)
}

// Repository struct that uses the ServiceConfig and implements RepositoryI
type Repository struct {
	Config *config.ServiceConfig
}

func (r *Repository) GetGridInfo(ctx context.Context, request models.Request) (string, error) {
	apiURL := fmt.Sprintf("%s/points/%s,%s", r.Config.URL, request.Latitude, request.Longitude)

	// Set the timeout based on the configuration
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.Config.Timeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("failed to call the upstream service: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("failed to read response body: %v", err)
		return "", err
	}

	var pointsRes PointsResponse
	if err := json.Unmarshal(body, &pointsRes); err != nil {
		log.Errorf("failed to unmarshal points response: %v", err)
		return "", err
	}

	return pointsRes.Properties.Forecast, nil
}

func (r *Repository) SearchWeatherApi(ctx context.Context, request models.Request) (WeatherResponse, error) {

	forecastURL, err := r.GetGridInfo(ctx, request)
	if err != nil {
		log.Error("failed to get grid info", err)
		return WeatherResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", forecastURL, nil)
	if err != nil {
		return WeatherResponse{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("failed to call the forecast service: %v", err)
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("failed to read forecast response body: %v", err)
		return WeatherResponse{}, err
	}

	// Unmarshal the weather data into WeatherResponse
	var weatherRes WeatherResponse
	if err := json.Unmarshal(body, &weatherRes); err != nil {
		log.Errorf("failed to unmarshal forecast response: %v", err)
		return WeatherResponse{}, err
	}

	return weatherRes, nil
}
