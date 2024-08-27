package mapper

import "weatherapi/v2/external/models"

type MapperI interface {
	MapWeatherApiResponse(res models.WeatherPropertiesRes) (response models.WeatherPropertiesRes, err error)
}

type Mapper struct {
}

func (m *Mapper) MapWeatherApiResponse(res models.WeatherPropertiesRes) (response models.WeatherPropertiesRes, err error) {
	//TODO implement me
	panic("implement me")
}
