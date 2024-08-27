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
	GetWeatherData(ctx context.Context, request models.Request) (res models.Response)
}

type Service struct {
	RepositoryService repository.RepositoryI
	Mapper            mapper.MapperI
}

func (r *Service) GetWeatherData(ctx context.Context, request models.Request) (res models.Response) {
	// Validate request
	if err := validateRequest(request); err != nil {
		log.Errorf("validate request error: %v", err)

		// Return the response with the error log
		return models.Response{
			Message: models.Message{
				ErrorLog: models.ErrorLogs{
					{
						Scope:      "Bad Request",
						StatusCode: "400", // Use 400 for Bad Request
						RootCause:  err.Error(),
					},
				},
			},
		}
	}

	// Call the repository method to fetch weather data
	apiData, err := r.RepositoryService.SearchWeatherApi(ctx, request)
	if err != nil {
		// Log the error
		log.Errorf("GetWeatherApiData error: %v", err)

		// Return the response with the error log
		return models.Response{
			Message: models.Message{
				ErrorLog: models.ErrorLogs{
					{
						Scope:      "Internal Server Error",
						StatusCode: "500",
						RootCause:  err.Error(),
					},
				},
			},
		}
	}

	// Perform mapping if necessary
	res = r.Mapper.MapWeatherData(apiData)

	return res
}

func validateRequest(request models.Request) error {

	// Parse latitude and longitude as float32
	latitude, err := strconv.ParseFloat(request.Latitude, 64)
	if err != nil {
		log.Printf("Invalid latitude: %v", err)
		return errors.New("invalid latitude")
	}

	longitude, err := strconv.ParseFloat(request.Longitude, 64)
	if err != nil {
		log.Printf("Invalid longitude: %v", err)
		return errors.New("invalid longitude")
	}

	// Validate latitude and longitude ranges
	if latitude < -90.0 || latitude > 90.0 {
		return errors.New("invalid latitude: must be between -90.0 and 90.0")
	}

	if longitude < -180.0 || longitude > 180.0 {
		return errors.New("invalid longitude: must be between -180.0 and 180.0")
	}

	return nil
}
