package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Boilerplate Fiber/GO")
	app := fiber.New()
	port := os.Getenv("PORT")
	app.Listen(port)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
