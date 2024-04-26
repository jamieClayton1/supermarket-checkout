package main

import (
	"supermarket-checkout/internal/api"
	"supermarket-checkout/internal/provider"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	log.Info("Generating service provider.")
	services := provider.NewApiServiceProvider()
	log.Info("Serving api.")
	api.NewAPI(&services).Serve()
}
