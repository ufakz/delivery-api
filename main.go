package main

import (
	"net/http"

	"github.com/farouqu/delivery-api/delivery"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/delivery", delivery.GetDeliveryEstimate).Methods("GET")
	router.HandleFunc("/delivery", delivery.CreateNewDelivery).Methods("POST")

	return router
}

func main() {
	r := Router()

	http.ListenAndServe(":8080", r)
}
