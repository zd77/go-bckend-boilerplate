package main

import (
	"fmt"
	"go-boilerplate/api"
	"go-boilerplate/config"
	"go-boilerplate/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Boilerplate net/http for GO")

	cfg := config.LoadConfig()

	_, err := db.ConnectMongoDb(cfg.MongoUrl)
	if err != nil {
		panic(err)
	}

	_, err = db.ConnectPostgresDb(cfg.PostgresUrl)
	if err != nil {
		panic(err)
	}

	router := api.RouterApi()

	port := os.Getenv("PORT")
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(port, router)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
