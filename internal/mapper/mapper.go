package mapper

import "weatherapi/v2/external/models"

type MapperI interface {
	MapWeatherData(res models.WeatherPropertiesRes) (response models.Response)
}

type Mapper struct {
}

func (m *Mapper) MapWeatherData(res models.WeatherPropertiesRes) (response models.Response) {
	//TODO implement me
	panic("implement me")
}
