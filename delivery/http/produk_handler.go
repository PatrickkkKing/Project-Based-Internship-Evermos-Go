package http

import (
	"evermos/domain"
	"evermos/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProdukHandler struct {
	uc *usecase.ProdukUsecase
}

func NewProdukHandler(uc *usecase.ProdukUsecase) *ProdukHandler {
	return &ProdukHandler{uc}
}

func (h *ProdukHandler) Create(c *fiber.Ctx) error {
	tid := c.Locals("toko_id")
	tokoID, ok := tid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	p := new(domain.Produk)
	if err := c.BodyParser(p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	file, err := c.FormFile("image")
	if err == nil {
		path := "./uploads/" + file.Filename
		_ = c.SaveFile(file, path)
		p.Image = path
	}

	if err := h.uc.Create(tokoID, p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "produk created"})
}

func (h *ProdukHandler) List(c *fiber.Ctx) error {
	tid := c.Locals("toko_id")
	tokoID, ok := tid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	list, err := h.uc.List(tokoID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(list)
}

func (h *ProdukHandler) Update(c *fiber.Ctx) error {
	tid := c.Locals("toko_id")
	tokoID, ok := tid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	p := new(domain.Produk)
	if err := c.BodyParser(p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	file, err := c.FormFile("image")
	if err == nil {
		path := "./uploads/" + file.Filename
		_ = c.SaveFile(file, path)
		p.Image = path
	}

	if err := h.uc.Update(tokoID, uint(id), p); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "produk updated"})
}

func (h *ProdukHandler) Delete(c *fiber.Ctx) error {
	tid := c.Locals("toko_id")
	tokoID, ok := tid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := h.uc.Delete(tokoID, uint(id)); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "produk deleted"})
}
