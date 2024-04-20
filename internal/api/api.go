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

type API struct {
	ServiceProvider *provider.ServiceProvider
}

func NewAPI(serviceProvider *provider.ServiceProvider) *API {
	return &API{
		ServiceProvider: serviceProvider,
	}
}

func (api API) Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/checkout/price", FetchCheckoutPriceHandler(api.ServiceProvider)).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}

func FetchCheckoutPriceHandler(provider *provider.ServiceProvider) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		body, err := UnpackBody[entity.FetchCheckoutPriceRequest](res, req)
		if err != nil {
			SendError(res, http.StatusBadRequest, err.Error())
			return
		}
		if len(body.ItemSKUs) == 0 {
			SendError(res, http.StatusBadRequest, "provide a list of item SKUs with a length greater than 0")
			return
		}
		result, err := provider.CheckoutService.FetchPrice(&entity.FetchPriceConfig{
			ItemSKUs: body.ItemSKUs,
		})
		if err != nil {
			SendError(res, http.StatusInternalServerError, err.Error())
			return
		}
		Send(res, map[string]any{
			"price": result.Price,
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
