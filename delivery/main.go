package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func GetDeliveryEstimate(w http.ResponseWriter, r *http.Request) {
	var m map[string]interface{}
	//Decode request body as JSON
	json.NewDecoder(r.Body).Decode(&m)

	//Prepend API KEY to request body
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
	var m map[string]interface{}
	//Decode request body as JSON
	json.NewDecoder(r.Body).Decode(&m)

	//Get firestore client
	ctx, client := InitializeFirestore()

	defer client.Close()

	//Save order information in firestore
	_, _, err := client.Collection("orders").Add(*ctx, m)

	if err != nil {
		log.Fatalf("Failed adding document: %v", err)
	}

	//Prepend API KEY to request body
	m["api_key"] = os.Getenv("API_KEY")
	requestBody, err := json.Marshal(m)
	if err != nil {
		return
	}

	//Make post request to Gokada API
	resp, err := http.Post("https://api.gokada.ng/api/developer/order_create", "application/json", bytes.NewBuffer(requestBody))
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
	var response DeliveryResponse
	json.Unmarshal(bodyBytes, &response)

	//Store order response in firestore
	_, _, err = client.Collection("delivery_orders").Add(*ctx, response)

	if err != nil {
		log.Fatalf("Failed adding document: %v", err)
	}

	//Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//Encode response as JSON
	json.NewEncoder(w).Encode(response)

}

func InitializeFirestore() (*context.Context, *firestore.Client) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "delivery-api-3cead"}
	sa := option.WithCredentialsFile("delivery-api.json")
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, nil
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &ctx, client
}
