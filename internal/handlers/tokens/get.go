package tokens

import (
	"context"
	"github.com/Ambassador4ik/medods-test-go/internal/jwt"
	"github.com/Ambassador4ik/medods-test-go/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

func getTokens(c *fiber.Ctx) error {
	userGUID := c.Query("guid")

	err := validator.New().Var(userGUID, "required,uuid4")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid GUID"})
	}

	parsedGUID, err := uuid.Parse(userGUID)

	ip := c.IP()

	accessTokenId := uuid.New()
	accessToken, err := jwt.GenerateAccessToken(parsedGUID, ip, accessTokenId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate access token"})
	}

	refreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate access token"})
	}

	refreshTokenHash, err := jwt.HashRefreshToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash refresh token"})
	}

	_, err = dbclient.Client.Token.Create().
		SetUserID(parsedGUID).
		SetToken(refreshTokenHash).
		SetAccessTokenID(accessTokenId).
		Save(context.Background())
	if err != nil {
		log.Print(err)
		// TODO: Differentiate userNotFound and other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to store refresh token"})
	}

	return c.JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
