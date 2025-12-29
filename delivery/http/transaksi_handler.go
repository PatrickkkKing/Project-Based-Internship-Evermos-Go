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

// ================= CREATE TRANSAKSI =================
func (h *TransaksiHandler) Create(c *fiber.Ctx) error {
	// ambil user_id aman
	uid := c.Locals("user_id")
	userID, ok := uid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	// ambil toko_id aman
	tid := c.Locals("toko_id")
	tokoID, ok := tid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	var req struct {
		AlamatID uint                   `json:"alamat_id"`
		Items    []domain.TransaksiItem `json:"items"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := h.uc.Create(userID, tokoID, req.AlamatID, req.Items); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "transaksi berhasil",
	})
}

// ================= LIST TRANSAKSI + PAGINATION =================
func (h *TransaksiHandler) List(c *fiber.Ctx) error {
	uid := c.Locals("user_id")
	userID, ok := uid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

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

// ================= FILTER TRANSAKSI BY DATE =================
func (h *TransaksiHandler) Filter(c *fiber.Ctx) error {
	uid := c.Locals("user_id")
	userID, ok := uid.(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	start := c.Query("start")
	end := c.Query("end")

	if start == "" || end == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "start and end date required",
		})
	}

	data, err := h.uc.FilterByDate(userID, start, end)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}
