package repository

import (
	"evermos/domain"

	"gorm.io/gorm"
)

type produkMySQL struct {
	db *gorm.DB
}

func NewProdukMySQL(db *gorm.DB) ProdukRepository {
	return &produkMySQL{db}
}

func (r *produkMySQL) Create(p *domain.Produk) error {
	return r.db.Create(p).Error
}

func (r *produkMySQL) List(tokoID uint) ([]domain.Produk, error) {
	var list []domain.Produk
	err := r.db.Where("toko_id = ?", tokoID).Find(&list).Error
	return list, err
}

func (r *produkMySQL) FindByID(id uint) (*domain.Produk, error) {
	var p domain.Produk
	err := r.db.First(&p, id).Error
	return &p, err
}

func (r *produkMySQL) Update(p *domain.Produk) error {
	return r.db.Save(p).Error
}

func (r *produkMySQL) Delete(id uint) error {
	return r.db.Delete(&domain.Produk{}, id).Error
}
