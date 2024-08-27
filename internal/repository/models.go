package repository

// Root object to capture the forecast URL from the points API response
type PointsResponse struct {
	Properties PointsProperties `json:"properties"`
}

// Struct for the "properties" field in the points API response
type PointsProperties struct {
	Forecast string `json:"forecast"`
}

// Root object for the response from the weather API
type WeatherResponse struct {
	Type       string             `json:"type"`
	Geometry   Geometry           `json:"geometry"`
	Properties ForecastProperties `json:"properties"`
}

// Struct for the "geometry" field in the response
type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// Struct for the "properties" field in the forecast response
type ForecastProperties struct {
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       string    `json:"generatedAt"`
	UpdateTime        string    `json:"updateTime"`
	ValidTimes        string    `json:"validTimes"`
	Elevation         Elevation `json:"elevation"`
	Periods           []Period  `json:"periods"` // Array of periods
}

// Struct for the "elevation" field in the "forecast properties"
type Elevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

// Struct for each "period" in the "forecast properties.periods"
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

// Struct for the "probabilityOfPrecipitation" field in the "periods"
type PoP struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}
