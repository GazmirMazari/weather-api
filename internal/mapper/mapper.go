package mapper

import (
	"weatherapi/v2/external/models"
	"weatherapi/v2/internal/repository"
)

type MapperI interface {
	MapWeatherData(res repository.WeatherResponse) (response models.Response)
}

type Mapper struct {
}

func (m *Mapper) MapWeatherData(res repository.WeatherResponse) (response models.Response) {

	// Iterate over the weather periods and look for "Today"
	for _, period := range res.Properties.Periods {
		if period.Name == "Today" {
			weatherInfo := models.Periods{
				Name:          period.Name,
				Temperature:   period.Temperature,
				ShortForecast: period.ShortForecast,
				Description:   classifyTemperature(period.Temperature),
			}
			response.WeatherResponse = append(response.WeatherResponse, weatherInfo)
			break
		}
	}

	// Return the mapped response
	return response
}

func classifyTemperature(temperature int) string {
	switch {
	case temperature >= 85:
		return "Hot"
	case temperature <= 60:
		return "Cold"
	default:
		return "Moderate"
	}
}
