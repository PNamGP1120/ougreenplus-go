package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
)

func JWTProtected(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return fiber.ErrUnauthorized
		}

		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			return fiber.ErrUnauthorized
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Locals("user_id", uint(claims["sub"].(float64)))
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
