package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/farouqu/delivery-api/delivery"
	_ "github.com/farouqu/delivery-api/docs"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	//Load environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	fmt.Println("Environment variables successfully loaded. Starting application...")

	r := delivery.Router()
	port, _ := os.LookupEnv("PORT")
	if port == "" {
		port = "8000"
	}

	//Allow cross-origin request for external API calls
	c := cors.AllowAll()

	//Configure http server with cors
	server := &http.Server{
		Handler:      c.Handler(r),
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	//Start server
	fmt.Println("Delivery API running on port ", port)
	log.Fatal(server.ListenAndServe())
}
