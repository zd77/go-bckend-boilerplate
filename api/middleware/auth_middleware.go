package middleware

import (
	"go-boilerplate/jwt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.GetReqHeaders()["Authorization"]
	tokenString := token[0]
	claims, err := jwt.VerifyToken(tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Message": "unauthorized request"})
	}
	expiresAt := int64(claims["ExpiresAt"].(float64))
	if expiresAt < time.Now().Unix() {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Message": "token expired"})
	}
	return c.Next()
}
