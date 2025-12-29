package repository

import "evermos/domain"

type TransaksiRepository interface {
	Create(tx *domain.Transaksi) error
	ListByUser(userID uint) ([]domain.Transaksi, error)
	ListByUserPaginate(userID uint, page, limit int) ([]domain.Transaksi, int64, error)
	FilterByDate(userID uint, start, end string) ([]domain.Transaksi, error)
}
