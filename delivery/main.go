package delivery

import (
	"encoding/json"
	"net/http"
)

func GetDeliveryEstimate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal("{'data':'Estimate Returned'}")
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func CreateNewDelivery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal("{'data':'Delivery Posted'}")
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
