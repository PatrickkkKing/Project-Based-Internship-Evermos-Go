package http

import (
	"evermos/usecase"

	"github.com/gofiber/fiber/v2"
)

type TokoHandler struct {
	tokoUC *usecase.TokoUsecase
}

func NewTokoHandler(tokoUC *usecase.TokoUsecase) *TokoHandler {
	return &TokoHandler{tokoUC}
}

func (h *TokoHandler) MyToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	toko, err := h.tokoUC.GetMyToko(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "toko not found"})
	}

	return c.JSON(toko)
}

func (h *TokoHandler) UpdateMyToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.tokoUC.UpdateMyToko(userID, req.Name); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "toko updated"})
}
