package tokens

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func refreshTokens(c *fiber.Ctx) error {
	accessToken := c.Query("accessToken")
	refreshToken := c.Query("refreshToken")

	errA := validator.New().Var(accessToken, "required,jwt")
	if errA != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid access token"})
	}

	errR := validator.New().Var(refreshToken, "required,base64")
	if errR != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid refresh token"})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid refresh token"})
}
