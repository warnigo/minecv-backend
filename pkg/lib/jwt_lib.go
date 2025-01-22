package lib

import (
	"time"

	"minecv/config"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims represents the JWT claims
type CustomClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new access token
func GenerateToken(userId string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.AppConfig.SecretKey
	return token.SignedString([]byte(secret))
}

// GenerateRefreshToken creates a new refresh token
func GenerateRefreshToken(userId string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.AppConfig.SecretKey
	return token.SignedString([]byte(secret))
}

// ParseToken validates the JWT token and returns claims
func ParseToken(tokenString string) (*CustomClaims, error) {
	secret := config.AppConfig.SecretKey
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
