package facade

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"strconv"
	"weatherapi/v2/external/models"
	"weatherapi/v2/internal/mapper"
	"weatherapi/v2/internal/repository"
)

type ServiceI interface {
	GetWeatherData(ctx context.Context, request models.Request) (response models.WeatherPropertiesRes, err error)
}

type Service struct {
	RepositoryService repository.RepositoryI
	Mapper            mapper.MapperI
}

func (r *Service) GetWeatherData(ctx context.Context, request models.Request) (res models.WeatherPropertiesRes, err error) {
	// Validate request
	if err := validateRequest(request); err != nil {
		log.Errorf("validate request error: %v", err)
		return res, err
	}

	// Call the repository method to fetch weather data
	response, err := r.RepositoryService.SearchWeatherApi(ctx, request)
	if err != nil {

		log.Errorf("GetWeatherApiData error: %v", err)
		return res, err
	}

	// Perform mapping if necessary
	// response = r.Mapper.MapWeatherData(response)

	return response, err
}

func validateRequest(request models.Request) error {
	// Check if latitude or longitude is empty
	if request.Latitude == "" || request.Longitude == "" {
		return errors.New("invalid request: latitude and longitude should not be empty")
	}

	// Convert latitude and longitude from strings to float64
	lat, err := strconv.ParseFloat(request.Latitude, 64)
	if err != nil {
		return errors.New("invalid latitude: must be a valid number")
	}

	lon, err := strconv.ParseFloat(request.Longitude, 64)
	if err != nil {
		return errors.New("invalid longitude: must be a valid number")
	}

	// Validate latitude and longitude ranges
	if lat < -90.0 || lat > 90.0 {
		return errors.New("invalid latitude: must be between -90.0 and 90.0")
	}

	if lon < -180.0 || lon > 180.0 {
		return errors.New("invalid longitude: must be between -180.0 and 180.0")
	}

	return nil
}
