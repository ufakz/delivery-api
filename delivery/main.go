package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetDeliveryEstimate(w http.ResponseWriter, r *http.Request) {
	//Encode request body as JSON
	reqBody, err := json.Marshal(r.Body)
	if err != nil {
		return
	}

	//Prepend API KEY to request body
	var m map[string]interface{}
	json.Unmarshal(reqBody, &m)
	m["api_key"] = os.Getenv("API_KEY")
	requestBody, err := json.Marshal(m)
	if err != nil {
		return
	}

	//Make post request to Gokada API
	resp, err := http.Post("https://api.gokada.ng/api/developer/order_estimate", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	//Read Gokada API response and store data
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v", string(bodyBytes))
	if err != nil {

		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	//Convert response to EstimateResponse struct
	var response EstimateResponse
	json.Unmarshal(bodyBytes, &response)

	//Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//Encode response as JSON
	json.NewEncoder(w).Encode(response)
}

func CreateNewDelivery(w http.ResponseWriter, r *http.Request) {

}
