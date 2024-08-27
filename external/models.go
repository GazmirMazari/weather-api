package external

type Request struct {
	Latitude  string `json:"city,omitempty"`
	Longitude string `json:"country,omitempty"`
}

type WeatherPropertiesResponse struct {
	Periods []struct {
		Name          string `json:"name"`
		Temperature   int    `json:"temperature"`
		ShortForecast string `json:"shortForecast"`
	} `json:"periods"`
}
