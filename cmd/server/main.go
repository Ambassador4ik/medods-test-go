package main

import (
	"github.com/Ambassador4ik/medods-test-go/internal/config"
	"github.com/Ambassador4ik/medods-test-go/internal/handlers/tokens"
	"github.com/Ambassador4ik/medods-test-go/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()
	dbclient.InitEntClient(cfg)

	tokens.RegisterRoutes(app)

	app.Listen(":3000")
}
