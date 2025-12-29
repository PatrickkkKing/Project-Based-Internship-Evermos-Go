package http

import (
	"evermos/domain"
	"evermos/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type KategoriHandler struct {
	uc *usecase.KategoriUsecase
}

func NewKategoriHandler(uc *usecase.KategoriUsecase) *KategoriHandler {
	return &KategoriHandler{uc}
}

func (h *KategoriHandler) Create(c *fiber.Ctx) error {
	role := c.Locals("role").(string)

	var req domain.Kategori
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := h.uc.Create(role, &req); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "kategori created"})
}

func (h *KategoriHandler) List(c *fiber.Ctx) error {
	list, _ := h.uc.List()
	return c.JSON(list)
}

func (h *KategoriHandler) Update(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	id, _ := strconv.Atoi(c.Params("id"))

	var req domain.Kategori
	c.BodyParser(&req)

	if err := h.uc.Update(role, uint(id), &req); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "kategori updated"})
}

func (h *KategoriHandler) Delete(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.uc.Delete(role, uint(id)); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "kategori deleted"})
}
