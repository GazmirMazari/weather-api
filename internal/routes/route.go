package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"weatherapi/v2/internal/facade"
	"weatherapi/v2/internal/models"
)

type Handler struct {
	Service *facade.Service
}

func (h *Handler) InitializeRoutes(opts ...options) {
	r := mux.NewRouter()

	// Health check
	r.Handle(http.MethodGet, "/health", h.HealthCheck())
	// Get forecast handler
	r.Handle(http.MethodGet, "/forecast", h.GetForecast())
}

func (h *Handler) GetForecast() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sw := time.Now()
		statusCode := http.StatusOK

		var apiRequest models.Request
		var apiResponse models.WeatherPropertiesResponse

		apiResponse = h.Service.GetWeatherData(ctx, *apiRequest.FromJSON(ctx.Request.Body))

		if len(apiResponse.Message.ErrorLog) > 0 {
			statusCode = apiResponse.Message.ErrorLog.GetHTTPStatus(len(apiResponse.Stuff))
		}
		apiResponse.Message.AddMessageDetails(sw)

		ctx.JSON(statusCode, apiResponse)
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
