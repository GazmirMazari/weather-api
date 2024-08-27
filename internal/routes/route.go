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

		statusCode := http.StatusOK

		apiRequest := queryParams(r)
		var apiResponse models.Response

		apiResponse = h.Service.GetWeatherData(ctx, apiRequest)

		//set the api error here
		if len(apiResponse.Message.ErrorLog) > 0 {
			statusCode = apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.WeatherResponse))
		}

		w.Header().Set("Content-Type", "application/json")
		apiResponse.Message.AddMessageDetails(sw)
		w.WriteHeader(statusCode)

		//Encode the response as json and writes to the response writer
		if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
	latitude := r.URL.Query().Get("latitude")

	longitude := r.URL.Query().Get("longitude")

	return models.Request{
		Latitude:  latitude,
		Longitude: longitude,
	}

}
