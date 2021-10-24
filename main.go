package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/farouqu/delivery-api/delivery"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/delivery_estimate", delivery.GetDeliveryEstimate).Methods(http.MethodPost)
	router.HandleFunc("/dellivery_order", delivery.CreateNewDelivery).Methods(http.MethodPost)

	return router
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	fmt.Println("Environment variables successfully loaded. Starting application...")

	r := Router()
	port, _ := os.LookupEnv("PORT")
	if port == "" {
		port = "8000"
	}

	c := cors.AllowAll()

	server := &http.Server{
		Handler:      c.Handler(r),
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Println("Delivery API running on port ", port)
	log.Fatal(server.ListenAndServe())
}
