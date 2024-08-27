package routes

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"weatherapi/v2/external/models"
	"weatherapi/v2/internal/facade"
)

type Handler struct {
	Service *facade.Service
}

func (h *Handler) InitializeRoutes(ctx context.Context) http.Handler {
	r := mux.NewRouter()

	// Health check
	r.Handle("/health", h.HealthCheck()).Methods(http.MethodGet)
	// Get forecast handler
	r.Handle("/forecast", h.GetForecast(ctx)).Methods(http.MethodGet)

	return r
}

func (h *Handler) GetForecast(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sw := time.Now()

		var statusCode int
		var apiError error
		apiRequest := queryParams(r)
		var apiResponse models.Response

		res, apiError := h.Service.GetWeatherData(ctx, apiRequest)
		if apiError != nil {
			log.Println(apiError)
		}

		if len(apiResponse.Message.ErrorLog) > 0 {
			statusCode = apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.WeatherPropertiesRes.Periods))
		}
		apiResponse.Message.AddMessageDetails(sw)
		apiResponse.WeatherPropertiesRes = res
		w.WriteHeader(statusCode)
	}
}

func (h *Handler) HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func queryParams(r *http.Request) models.Request {
	latitude := r.URL.Query()["latitude"]
	longitude := r.URL.Query()["longitude"]

	return models.Request{
		Latitude:  latitude[0],
		Longitude: longitude[0],
	}

}
