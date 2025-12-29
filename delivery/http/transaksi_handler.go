package http

import (
	"evermos/domain"
	"evermos/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransaksiHandler struct {
	uc *usecase.TransaksiUsecase
}

func NewTransaksiHandler(uc *usecase.TransaksiUsecase) *TransaksiHandler {
	return &TransaksiHandler{uc}
}

func (h *TransaksiHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	tokoID := c.Locals("toko_id").(uint)

	var req struct {
		AlamatID uint                   `json:"alamat_id"`
		Items    []domain.TransaksiItem `json:"items"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.uc.Create(userID, tokoID, req.AlamatID, req.Items)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "transaksi berhasil"})
}

// func (h *TransaksiHandler) List(c *fiber.Ctx) error {
// 	userID := c.Locals("user_id").(uint)

// 	list, err := h.uc.ListByUser(userID)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(list)
// }

func (h *TransaksiHandler) List(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	data, total, err := h.uc.List(userID, page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func (h *TransaksiHandler) Filter(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	start := c.Query("start")
	end := c.Query("end")

	data, err := h.uc.FilterByDate(userID, start, end)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}
