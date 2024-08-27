package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Config *config.Config
}

func (r *Repository) SearchWeatherApi(ctx context.Context, request Request) (response Response, err error) {
	//build the request
	requestUrl := fmt.Sprintf("https://api.weather.gov/points/%s,%s", request.Latitude, request.Longitude)

	res, err := http.Get(requestUrl)
	if err != nil {
		return response, fmt.Errorf("GetWeatherApiData error: %v", err)
	}

	defer res.Body.Close()

	//unmarshal the response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("failed to unmarshall the error: %v", err)
	}

	return response, nil
}
