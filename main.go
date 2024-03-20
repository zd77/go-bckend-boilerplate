package main

import (
	"fmt"
	"go-boilerplate/api"
	"go-boilerplate/config"
	"go-boilerplate/db"
	"go-boilerplate/stores"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Boilerplate net/http for GO")

	cfg := config.LoadConfig()
	app := fiber.New()

	mongoClient, err := db.ConnectMongoDb(cfg.MongoUrl)
	if err != nil {
		panic(err)
	}

	// _, err = db.ConnectPostgresDb(cfg.PostgresUrl)
	// if err != nil {
	// 	panic(err)
	// }

	routerDeps := api.RouterDependencies{
		TaskStore: stores.NewTaskStore(mongoClient),
		UserStore: stores.NewUserStore(mongoClient),
	}

	api.RouterApi(app, routerDeps)

	port := os.Getenv("PORT")
	fmt.Printf("Server listening on port %s\n", port)
	app.Listen(port)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
