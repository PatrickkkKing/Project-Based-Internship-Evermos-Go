package repository

import "evermos/domain"

type KategoriRepository interface {
	Create(kategori *domain.Kategori) error
	List() ([]domain.Kategori, error)
	FindByID(id uint) (*domain.Kategori, error)
	Update(kategori *domain.Kategori) error
	Delete(id uint) error
}
