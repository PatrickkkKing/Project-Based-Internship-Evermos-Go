package domain

import "time"

type Produk struct {
	ID        uint `gorm:"primaryKey"`
	TokoID    uint
	Name      string
	Desc      string
	Price     float64
	Stock     int
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
