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
	tokoID := c.Locals("toko_id").(uint)

	p := new(domain.Produk)
	if err := c.BodyParser(p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	file, err := c.FormFile("image")
	if err == nil {
		path := "./uploads/" + file.Filename
		c.SaveFile(file, path)
		p.Image = path
	}

	if err := h.uc.Create(tokoID, p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "produk created"})
}

func (h *ProdukHandler) List(c *fiber.Ctx) error {
	tokoIDRaw := c.Locals("toko_id")
	if tokoIDRaw == nil {
		return c.Status(400).JSON(fiber.Map{"error": "toko_id not found"})
	}
	tokoID := tokoIDRaw.(uint)
	list, _ := h.uc.List(tokoID)
	return c.JSON(list)
}

func (h *ProdukHandler) Update(c *fiber.Ctx) error {
	tokoID := c.Locals("toko_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	p := new(domain.Produk)
	c.BodyParser(p)

	file, err := c.FormFile("image")
	if err == nil {
		path := "./uploads/" + file.Filename
		c.SaveFile(file, path)
		p.Image = path
	}

	if err := h.uc.Update(tokoID, uint(id), p); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "produk updated"})
}

func (h *ProdukHandler) Delete(c *fiber.Ctx) error {
	tokoID := c.Locals("toko_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.uc.Delete(tokoID, uint(id)); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "produk deleted"})
}
