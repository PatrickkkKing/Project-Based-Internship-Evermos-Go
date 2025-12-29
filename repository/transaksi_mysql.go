package repository

import (
	"evermos/domain"

	"gorm.io/gorm"
)

type transaksiMySQL struct {
	db *gorm.DB
}

func NewTransaksiMySQL(db *gorm.DB) TransaksiRepository {
	return &transaksiMySQL{db}
}

func (r *transaksiMySQL) Create(tx *domain.Transaksi) error {
	return r.db.Create(tx).Error
}

func (r *transaksiMySQL) ListByUser(userID uint) ([]domain.Transaksi, error) {
	var list []domain.Transaksi
	err := r.db.Preload("Items").
		Where("user_id = ?", userID).
		Find(&list).Error
	return list, err
}

func (r *transaksiMySQL) ListByUserPaginate(
	userID uint,
	page, limit int,
) ([]domain.Transaksi, int64, error) {

	var list []domain.Transaksi
	var total int64

	offset := (page - 1) * limit

	r.db.Model(&domain.Transaksi{}).
		Where("user_id = ?", userID).
		Count(&total)

	err := r.db.
		Preload("Items").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&list).Error

	return list, total, err
}

func (r *transaksiMySQL) FilterByDate(
	userID uint,
	start, end string,
) ([]domain.Transaksi, error) {

	var list []domain.Transaksi

	err := r.db.
		Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, start, end).
		Preload("Items").
		Find(&list).Error

	return list, err
}
