package jwt

import (
	"time"

	"github.com/ahsansaif47/home-kitchens/auth/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(config.GetConfig().JWTSecret)

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

// I wont be needing this in the auth service but this will be used in other services
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("email", claims.Email)
		c.Locals("user_name", claims.UserName)
		c.Locals("role_id", claims.Role)

		return c.Next()
	}
}
