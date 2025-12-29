package repository

import "evermos/domain"

type LogProdukRepository interface {
	Create(log *domain.LogProduk) error
}
