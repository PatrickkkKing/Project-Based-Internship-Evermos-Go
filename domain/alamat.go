package domain

import "time"

type Alamat struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Label     string
	Alamat    string
	Kota      string
	Provinsi  string
	KodePos   string
	CreatedAt time.Time
}
