package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email    string
	UserName string
	Role     uint
	jwt.RegisteredClaims
}

func GenerateJWT(email, userName string, roleID uint, secretKey string) (string, error) {
	expTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email:    email,
		UserName: userName,
		Role:     roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "home-kitchens",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
