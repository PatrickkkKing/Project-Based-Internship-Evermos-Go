package repository

import (
	"evermos/domain"

	"gorm.io/gorm"
)

type kategoriMySQL struct {
	db *gorm.DB
}

func NewKategoriMySQL(db *gorm.DB) KategoriRepository {
	return &kategoriMySQL{db}
}

func (r *kategoriMySQL) Create(kategori *domain.Kategori) error {
	return r.db.Create(kategori).Error
}

func (r *kategoriMySQL) List() ([]domain.Kategori, error) {
	var list []domain.Kategori
	err := r.db.Find(&list).Error
	return list, err
}

func (r *kategoriMySQL) FindByID(id uint) (*domain.Kategori, error) {
	var k domain.Kategori
	err := r.db.First(&k, id).Error
	return &k, err
}

func (r *kategoriMySQL) Update(kategori *domain.Kategori) error {
	return r.db.Save(kategori).Error
}

func (r *kategoriMySQL) Delete(id uint) error {
	return r.db.Delete(&domain.Kategori{}, id).Error
}
