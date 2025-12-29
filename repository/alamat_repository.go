package repository

import "evermos/domain"

type AlamatRepository interface {
	Create(alamat *domain.Alamat) error
	FindByUserID(userID uint) ([]domain.Alamat, error)
	FindByID(id uint) (*domain.Alamat, error)
	Update(alamat *domain.Alamat) error
	Delete(id uint) error
}
