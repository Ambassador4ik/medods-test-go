package jwt

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Ambassador4ik/medods-test-go/ent/token"
	dbclient "github.com/Ambassador4ik/medods-test-go/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// TODO: Get from env
const clientSecret = "dfss8sdfhsdfsd98f"

type CustomClaims struct {
	ID   uuid.UUID `json:"id"`
	GUID uuid.UUID `json:"guid"`
	IP   string    `json:"ip"`
	Exp  int64     `json:"exp"`
	jwt.RegisteredClaims
}

func (c CustomClaims) Valid() error {
	if c.Exp < time.Now().Unix() {
		return errors.New("token has expired")
	}
	return nil
}

func GenerateAccessToken(userGUID uuid.UUID, ip string, tokenId uuid.UUID) (string, error) {
	claims := CustomClaims{
		ID:   tokenId,
		GUID: userGUID,
		IP:   ip,
		Exp:  time.Now().Add(time.Hour * 1).Unix(),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := newToken.SignedString([]byte(clientSecret))
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

func ParseAccessToken(tokenStr string) (*CustomClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(clientSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*CustomClaims)
	if !ok || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateTokenPair(accessTokenClaims *CustomClaims, refreshTokenStr string) bool {
	tokenObj, err := dbclient.Client.Token.
		Query().
		Where(token.AccessTokenID(accessTokenClaims.ID)).
		Only(context.Background())
	if err != nil {
		return false
	}

	refreshTokenValid := bcrypt.CompareHashAndPassword([]byte(tokenObj.Token), []byte(refreshTokenStr))
	if refreshTokenValid != nil {
		return false
	}

	return true
}
