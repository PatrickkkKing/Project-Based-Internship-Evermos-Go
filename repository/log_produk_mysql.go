package repository

import (
	"evermos/domain"

	"gorm.io/gorm"
)

type LogProdukMySQL struct {
	db *gorm.DB
}

func NewLogProdukRepository(db *gorm.DB) *LogProdukMySQL {
	return &LogProdukMySQL{db}
}

func (r *LogProdukMySQL) Create(log *domain.LogProduk) error {
	return r.db.Create(log).Error
}
