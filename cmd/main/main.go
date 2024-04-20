package main

import (
	"supermarket-checkout/internal/api"
	"supermarket-checkout/internal/provider"
)

func main() {
	services := provider.NewApiServiceProvider()
	api.NewAPI(&services).Serve()
}
