package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/provider"

	"github.com/gorilla/mux"
)

// API containing a service provider
type API struct {
	ServiceProvider *provider.ServiceProvider
}

// Construct a new API, given a particular service provider
func NewAPI(serviceProvider *provider.ServiceProvider) *API {
	return &API{
		ServiceProvider: serviceProvider,
	}
}

// Serve the API on port 80
func (api API) Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/checkout/price", FetchCheckoutPriceHandler(api.ServiceProvider)).Methods("POST") // Use POST as pricing schema is not idempotent
	log.Fatal(http.ListenAndServe(":80", router)) // Customise behind an environment variable - for ease of use, we'll use default HTTP port
}

// Construct a handle for scanning an item, given a service provider
func ScanItemHandler(provider *provider.ServiceProvider) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		body, err := UnpackBody[entity.ScanItemRequest](res, req)
		if err != nil {
			SendError(res, http.StatusBadRequest, err.Error())
			return
		}
		if body.SKU == "" {
			SendError(res, http.StatusBadRequest, "provide a sku")
			return
		}
		basket_id, err := provider.CheckoutService.ScanItem(body.SKU, body.BasketId)
		if err != nil {
			SendError(res, http.StatusInternalServerError, err.Error())
			return
		}
		Send(res, map[string]any{
			"basket_id": basket_id,
		})
	}
}


// Construct a handle for fetching a checkout price, given a service provider
func FetchCheckoutPriceHandler(provider *provider.ServiceProvider) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		body, err := UnpackBody[entity.FetchCheckoutPriceRequest](res, req)
		if err != nil {
			SendError(res, http.StatusBadRequest, err.Error())
			return
		}
		if body.BasketId == "" {
			SendError(res, http.StatusBadRequest, "provide a basket_id")
			return
		}
		price, err := provider.CheckoutService.FetchPrice(body.BasketId)
		if err != nil {
			SendError(res, http.StatusInternalServerError, err.Error())
			return
		}
		Send(res, map[string]any{
			"price": price,
		})
	}
}

// Unpack request body from http.Request
// Returns either generic type T or http.StatusBadRequest error
func UnpackBody[T any](res http.ResponseWriter, req *http.Request) (T, error) {
	var container T
	b, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return container, err
	}
	err = json.Unmarshal(b, &container)
	defer req.Body.Close()
	return container, err
}

// Send json response via http.ResponseWriter
func Send(res http.ResponseWriter, payload map[string]any) {
	json.NewEncoder(res).Encode(payload)
}

// Send error with json response that has the following shape:
//
//	{
//		"error": message
//	}
func SendError(res http.ResponseWriter, code int, message string) {
	res.WriteHeader(code)
	json.NewEncoder(res).Encode(map[string]string{
		"error": message,
	})
}
