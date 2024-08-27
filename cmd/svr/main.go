package main

import (
	"context"
	"github.com/NYTimes/gziphandler"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"weatherapi/v2/cmd/svr/config"
	"weatherapi/v2/internal/routes"
)

const port = ":8080"
const configPath = "config/config.yaml"

func main() {
	defer recoverPanics()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//load the config
	cfg, err := config.New(configPath)
	if err != nil {
		log.Errorf("error loading config: %v", err)
	}
	//establish the facade
	service, err := InitUpstreamConfigClient(cfg)
	if err != nil {
		log.Errorf("error initializing upstream client: %v", err)
		os.Exit(1)
	}

	handler := routes.Handler{Service: service}

	router := handler.InitializeRoutes(ctx)

	log.Printf("Server is listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, gziphandler.GzipHandler(router)))

}

func recoverPanics() {
	if r := recover(); r != nil {
		log.Errorf("a panic has happened.: %v", r)
		log.Errorf("I should be alerting someone...")
	}
}
