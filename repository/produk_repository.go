package repository

import "evermos/domain"

type ProdukRepository interface {
	Create(produk *domain.Produk) error
	FindByID(id uint) (*domain.Produk, error)
	Update(produk *domain.Produk) error
	Delete(id uint) error
	List(tokoID uint) ([]domain.Produk, error)
}
