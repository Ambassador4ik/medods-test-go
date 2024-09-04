package tokens

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	tokenRoutes := app.Group("/tokens")
	tokenRoutes.Post("/get", getTokens)
	tokenRoutes.Post("/refresh", refreshTokens)
}
