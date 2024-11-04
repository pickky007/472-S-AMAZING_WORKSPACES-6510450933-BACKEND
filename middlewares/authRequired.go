package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// claim := token.Claims.(jwt.MapClaims)
	// fmt.Println(claim["username"])

	return c.Next()
}
