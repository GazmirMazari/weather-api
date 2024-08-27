package main

import (
	"context"
	log "github.com/sirupsen/logrus"
)

var config = "config/config.yaml"

func main() {
	//possible point of crash, recover in case of panics, function will execute in the end.
	defer recoverPanics()

	//establish a context which can be propagated
	ctx, cancel := context.WithCancel(context.Background())
	////load the config here
	appConfig := config.New(config)

	if service, err := initialize

	//
	////establish a facade here
	//
	//handler := routes.Handler

	//init the server

}

func recoverPanics() {
	if r := recover(); r != nil {
		log.Errof("paniced and am quitting &v", r)
		log.Errof("I should be alerting someone...")
	}
}
