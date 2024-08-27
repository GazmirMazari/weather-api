package mapper

import "weatherapi/v2/external/models"

type MapperI interface {
	MapWeatherData(res models.WeatherPropertiesRes) (response models.Response)
}

type Mapper struct {
}

func (m *Mapper) MapWeatherData(res models.WeatherPropertiesRes) (response models.Response) {
	// Initialize the response structure
	response = models.Response{}

	// Iterate over the weather periods and map the data to the response
	//for _, period := range res.Periods {
	//	// Extract relevant fields from the weather data
	//	//response.WeatherResponse = append(response.WeatherResponse, models.WeatherPropertiesRes{
	//	//	Name:          period.Name,
	//	//	Temperature:   period.Temperature,
	//	//	ShortForecast: period.ShortForecast,
	//	//})
	//}

	// Return the mapped response
	return response
}
