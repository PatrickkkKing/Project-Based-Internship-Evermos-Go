package middleware

import (
	"evermos/config"
	"evermos/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected(tokoRepo repository.TokoRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "missing token"})
		}

		// Format: Bearer <token>
		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))
		c.Locals("user_id", userID)
		c.Locals("role", claims["role"].(string))

		// Ambil toko dari userID
		toko, err := tokoRepo.FindByUserID(userID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "toko not found"})
		}
		c.Locals("toko_id", toko.ID)

		return c.Next()
	}
}
