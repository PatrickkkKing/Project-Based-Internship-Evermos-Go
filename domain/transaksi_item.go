package domain

type TransaksiItem struct {
	ID          uint `gorm:"primaryKey"`
	TransaksiID uint
	ProdukID    uint
	NamaProduk  string
	HargaProduk float64
	Qty         int
	Subtotal    float64
}
