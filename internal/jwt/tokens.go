package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pkcs12"
	"time"
)

// TODO: Get from env
const clientSecret = "dfss8sdfhsdfsd98f"

func GenerateAccessToken(userGUID string, ip string, tokenId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"id":   tokenId,
		"guid": userGUID,
		"ip":   ip,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(clientSecret))

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GenerateRefreshToken() (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(tokenBytes), nil
}

func HashRefreshToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidateAccessToken(token string) (string, error) {
	panic(pkcs12.NotImplementedError("df"))
}
