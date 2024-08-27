package repository

type PointsResponse struct {
	Properties PointsProperties `json:"properties"`
}

type PointsProperties struct {
	Forecast string `json:"forecast"`
}

type WeatherResponse struct {
	Type       string             `json:"type"`
	Geometry   Geometry           `json:"geometry"`
	Properties ForecastProperties `json:"properties"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type ForecastProperties struct {
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       string    `json:"generatedAt"`
	UpdateTime        string    `json:"updateTime"`
	ValidTimes        string    `json:"validTimes"`
	Elevation         Elevation `json:"elevation"`
	Periods           []Period  `json:"periods"`
}

type Elevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type Period struct {
	Number                     int    `json:"number"`
	Name                       string `json:"name"`
	StartTime                  string `json:"startTime"`
	EndTime                    string `json:"endTime"`
	IsDaytime                  bool   `json:"isDaytime"`
	Temperature                int    `json:"temperature"`
	TemperatureUnit            string `json:"temperatureUnit"`
	TemperatureTrend           string `json:"temperatureTrend"`
	ProbabilityOfPrecipitation PoP    `json:"probabilityOfPrecipitation"`
	WindSpeed                  string `json:"windSpeed"`
	WindDirection              string `json:"windDirection"`
	Icon                       string `json:"icon"`
	ShortForecast              string `json:"shortForecast"`
	DetailedForecast           string `json:"detailedForecast"`
}

type PoP struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}
