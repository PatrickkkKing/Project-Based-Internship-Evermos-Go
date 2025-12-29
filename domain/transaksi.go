package domain

import "time"

type Transaksi struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	TokoID    uint
	AlamatID  uint
	Total     float64
	CreatedAt time.Time

	Items []TransaksiItem `gorm:"foreignKey:TransaksiID"`
}
