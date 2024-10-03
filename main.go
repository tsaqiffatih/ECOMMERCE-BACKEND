package main

import (
	"log"
	"net/http"
	"os"
	_ "os"

	"github.com/tsaqiffatih/ecommerce-bakcend/config"
	"github.com/tsaqiffatih/ecommerce-bakcend/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	router := routers.InitRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	http.ListenAndServe(":"+port, router)
}
