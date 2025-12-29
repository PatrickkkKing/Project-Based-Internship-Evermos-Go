package domain

import "time"

type LogProduk struct {
	ID          uint `gorm:"primaryKey"`
	TransaksiID uint
	ProdukID    uint
	Nama        string
	Harga       float64
	Qty         int
	Subtotal    float64
	CreatedAt   time.Time
}
