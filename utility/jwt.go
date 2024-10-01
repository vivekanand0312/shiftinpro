package utility

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT utilities
type JWTClaim struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func GenerateJWT(phone string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claims := &JWTClaim{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 90).Unix(), // 90 days
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken(phone string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claims := &JWTClaim{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(), // 1 week
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateJWT(tokenString string) (*JWTClaim, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
