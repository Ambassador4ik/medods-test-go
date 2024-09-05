package tokens

import (
	"context"
	"github.com/Ambassador4ik/medods-test-go/ent/token"
	"github.com/Ambassador4ik/medods-test-go/internal/jwt"
	"github.com/Ambassador4ik/medods-test-go/internal/mail"
	dbclient "github.com/Ambassador4ik/medods-test-go/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
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

	// Decoded payload
	accessTokenClaims, err := jwt.ParseAccessToken(accessToken)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid access token"})
	}

	// Notify on IP change
	ip := c.IP()
	if ip != accessTokenClaims.IP {
		// Mocked email
		const email string = "user@example.com"
		mail.SendNewIpNotification(ip, email)
	}

	tokenPairValid := jwt.ValidateTokenPair(accessTokenClaims, refreshToken)
	if !tokenPairValid {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "invalid token pair"})
	}

	// Clean up database, ensuring one-time usage of refresh token
	_, err = dbclient.Client.Token.Delete().
		Where(token.AccessTokenID(accessTokenClaims.ID)).
		Exec(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not clean up database from old tokens"})
	}

	accessTokenId := uuid.New()
	newAccessToken, err := jwt.GenerateAccessToken(accessTokenClaims.GUID, ip, accessTokenId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate access token"})
	}

	newRefreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate access token"})
	}

	refreshTokenHash, err := jwt.HashRefreshToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash refresh token"})
	}

	// Stores updated session for the user
	_, err = dbclient.Client.Token.Create().
		SetUserID(accessTokenClaims.GUID).
		SetToken(refreshTokenHash).
		SetAccessTokenID(accessTokenId).
		Save(context.Background())
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to store refresh token"})
	}

	// We should probably consider 'set-cookie' instead
	return c.JSON(fiber.Map{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
