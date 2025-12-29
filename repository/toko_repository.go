package repository

import "evermos/domain"

type TokoRepository interface {
	Create(toko *domain.Toko) error
	FindByUserID(userID uint) (*domain.Toko, error)
	Update(toko *domain.Toko) error
}
