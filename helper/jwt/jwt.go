package jwt

import (
	"net/http"
	"time"

	"github.com/Reza1878/goesclearning/user-service/helper/fault"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const tokenExpiry = 30 * time.Minute
const refreshTokenExpiry = 72 * time.Hour

var signedKey = []byte("secret")

type JWTPayload struct {
	Name   string
	Email  string
	UserId string
	jwt.RegisteredClaims
}

func CreateAccessToken(name, email, userId string) (*string, *JWTPayload, error) {
	return generateToken(name, email, userId, tokenExpiry)
}

func CreateRefreshToken(name, email, userId string) (*string, *JWTPayload, error) {
	return generateToken(name, email, userId, refreshTokenExpiry)
}

func generateToken(name, email, userId string, duration time.Duration) (*string, *JWTPayload, error) {
	payload, err := newJWTPayload(name, email, userId, duration)
	if err != nil {
		return nil, nil, err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, payload).SignedString(signedKey)
	if err != nil {
		return nil, nil, fault.Custom(
			http.StatusUnprocessableEntity,
			fault.ErrUnprocessable,
			"failed signing JWT token: "+err.Error(),
		)
	}

	return &token, payload, nil
}

func newJWTPayload(name, email, userId string, duration time.Duration) (*JWTPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, fault.Custom(
			http.StatusUnprocessableEntity,
			fault.ErrUnprocessable,
			"failed to generate token ID: "+err.Error(),
		)
	}

	now := time.Now()
	exp := now.Add(duration)

	return &JWTPayload{
		Name:   name,
		Email:  email,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "user_login",
			Subject:   "go-escape",
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}, nil
}
