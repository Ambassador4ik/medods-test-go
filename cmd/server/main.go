package main

import (
	"github.com/Ambassador4ik/medods-test-go/internal/config"
	"github.com/Ambassador4ik/medods-test-go/internal/handlers/tokens"
	dbclient "github.com/Ambassador4ik/medods-test-go/internal/repository"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()
	dbclient.InitEntClient(cfg)

	tokens.RegisterRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Failed to start the server!")
	}
}
