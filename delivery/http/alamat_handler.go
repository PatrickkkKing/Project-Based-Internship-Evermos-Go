package http

import (
	"evermos/domain"
	"evermos/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AlamatHandler struct {
	uc *usecase.AlamatUsecase
}

func NewAlamatHandler(uc *usecase.AlamatUsecase) *AlamatHandler {
	return &AlamatHandler{uc}
}

func (h *AlamatHandler) Create(c *fiber.Ctx) error {
	uid := c.Locals("user_id")
	userID, ok := uid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	var req domain.Alamat
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := h.uc.Create(userID, &req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "alamat created"})
}

func (h *AlamatHandler) List(c *fiber.Ctx) error {
	uid := c.Locals("user_id")
	userID, ok := uid.(uint)
	if !ok {
		return c.SendStatus(401)
	}

	data, _ := h.uc.List(userID)
	return c.JSON(data)
}

func (h *AlamatHandler) Update(c *fiber.Ctx) error {
	uid := c.Locals("user_id")
	userID, ok := uid.(uint)
	if !ok {
		return c.SendStatus(401)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var req domain.Alamat
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := h.uc.Update(userID, uint(id), &req); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "alamat updated"})
}

func (h *AlamatHandler) Delete(c *fiber.Ctx) error {
	uid := c.Locals("user_id")
	userID, ok := uid.(uint)
	if !ok {
		return c.SendStatus(401)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := h.uc.Delete(userID, uint(id)); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "alamat deleted"})
}
