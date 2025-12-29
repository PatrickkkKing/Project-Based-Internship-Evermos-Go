package helper

import "github.com/gofiber/fiber/v2"

func GetUserID(c *fiber.Ctx) (uint, error) {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return 0, fiber.ErrUnauthorized
	}
	return userID, nil
}
