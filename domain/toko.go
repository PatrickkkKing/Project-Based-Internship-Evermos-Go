package domain

import "time"

type Toko struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Name      string
	CreatedAt time.Time
}
